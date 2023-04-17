package utility

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenUtility is a utility class to manage the issuance and validation of JWT tokens.
type TokenUtility struct {
	mutex               sync.RWMutex
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
	keyRotationInterval time.Duration
	keyRotationTicker   *time.Ticker
	expiration          *time.Time
}

// Singleton
var tokenUtilityInstance *TokenUtility
var tokenUtilityOnce sync.Once

func GetTokenUtilityInstance() (*TokenUtility, error) {
	keySize := 2048
	keyRotationInterval := 24 * time.Hour

	tokenUtilityOnce.Do(func() {
		var err error
		tokenUtilityInstance, err = newTokenUtility(keySize, keyRotationInterval)
		if err != nil {
			Log(ERROR, fmt.Sprintf("Failed to create token utility: %v", err))
		}
	})
	return tokenUtilityInstance, nil
}

func (tu *TokenUtility) GetPrivateKey() *rsa.PrivateKey {
	return tokenUtilityInstance.privateKey
}

// NewTokenUtility creates a new instance of TokenUtility.
// keySize: The size of the RSA key in bits (2048 or 4096 are recommended).
// keyRotationInterval: The interval between private key rotations.
func newTokenUtility(keySize int, keyRotationInterval time.Duration) (*TokenUtility, error) {
	// Generate the initial RSA private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		Log(ERROR, fmt.Sprintf("Failed to generate RSA key: %v", err))
		return nil, err
	}

	tu := &TokenUtility{
		privateKey:          privateKey,
		publicKey:           &privateKey.PublicKey,
		keyRotationInterval: keyRotationInterval,
	}

	// Start the key rotation.
	tu.keyRotationTicker = time.NewTicker(keyRotationInterval)
	go tu.rotateKey()

	Log(INFO, "Initialized TokenUtility")
	return tu, nil
}

// rotateKey periodically generates a new RSA private key and replaces the current one.
func (tu *TokenUtility) rotateKey() {
	for range tu.keyRotationTicker.C {
		newKey, err := rsa.GenerateKey(rand.Reader, tu.privateKey.Size()*8)
		if err != nil {
			Log(ERROR, fmt.Sprintf("Failed to rotate key: %v", err))

			continue
		}

		tu.mutex.Lock()
		tu.privateKey = newKey
		tu.publicKey = &newKey.PublicKey
		tu.mutex.Unlock()
		Log(INFO, "Rotated TokenUtility key successfully")
	}
}

// IssueToken generates a new JWT token with the provided claims and expiration time.
// claims: A map of custom claims that will be embedded in the token.
// expiresIn: The number of seconds before the token expires.
// Returns a signed JWT token.
func (tu *TokenUtility) IssueToken(customClaims map[string]interface{}) (string, string, error) {
	tu.mutex.RLock()
	defer tu.mutex.RUnlock()

	accessTokenDuration := time.Hour
	refreshTokenDuration := 24 * time.Hour

	// Create a new JWT token with the provided claims and expiration time.
	claims := jwt.MapClaims{
		"exp": time.Now().Add(accessTokenDuration).Unix(),
	}

	// Add custom claims.
	for k, v := range customClaims {
		claims[k] = v
	}

	// Sign the token using the RSA private key.
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedAccessToken, err := accessToken.SignedString(tu.privateKey)
	if err != nil {
		Log(ERROR, "Failed to sign the access token: %v", err)
		return "", "", err
	}

	// Create and sign the refresh token.
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	claims["exp"] = time.Now().Add(refreshTokenDuration).Unix() // Set the refresh token to expire later than the access token.
	signedRefreshToken, err := refreshToken.SignedString(tu.privateKey)
	if err != nil {
		Log(ERROR, "Failed to sign the refresh token: %v", err)
		return "", "", err
	}

	Log(INFO, "Issued token for user: %v", claims["sub"])
	return signedAccessToken, signedRefreshToken, nil
}

// ValidateToken parses and validates a JWT token using the RSA public key.
// tokenString: The JWT token string to parse and validate.
// Returns the parsed JWT token with its claims, or an error if the token is invalid or expired.
func (tu *TokenUtility) ValidateToken(tokenString string) (*jwt.Token, error) {
	tu.mutex.RLock()
	defer tu.mutex.RUnlock()

	// Parse the token and validate it using the RSA public key.
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Check if the token uses the expected signing method.
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			Log(ERROR, fmt.Sprintf("unexpected signing method: %v", t.Header["alg"]))
			return nil, errors.New("unexpected signing method")
		}

		return tu.publicKey, nil
	})

	if err != nil {
		Log(ERROR, fmt.Sprintf("Token is invalid or expired: %v", err))

		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	subject, _ := token.Claims.GetSubject()
	Log(INFO, "Validated token for user: %v", subject)
	return token, nil
}

func (tu *TokenUtility) RefreshToken(refreshTokenString string) (string, string, error) {
	// Check if the refresh token is still valid.
	refreshToken, err := tu.ValidateToken(refreshTokenString)
	if err != nil {
		// Refresh token is invalid or expired, deny access.
		Log(ERROR, fmt.Sprintf("Refresh token is invalid or expired: %v", err))
		return "", "", fmt.Errorf("refresh token is invalid or expired")
	}

	// Refresh token is valid, issue new access and refresh tokens.
	newAccessToken, newRefreshToken, err := tu.IssueToken(refreshToken.Claims.(jwt.MapClaims))
	if err != nil {
		Log(ERROR, fmt.Sprintf("Failed to generate new tokens: %v", err))
		return "", "", fmt.Errorf("failed to generate new tokens: %w", err)
	}
	return newAccessToken, newRefreshToken, nil
}
