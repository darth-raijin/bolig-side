package test

import (
	"testing"
	"time"

	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestCanIssueTokensWithClaims(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	claims := map[string]interface{}{
		"sub": "sample",
	}

	accessToken, refreshToken, err := tokenUtility.IssueToken(claims)

	// Check if there is no error
	assert.NoError(t, err)

	// Check if access and refresh tokens are not empty
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)

	// Validate the issued access token
	validatedToken, err := tokenUtility.ValidateToken(accessToken)

	// Check if there is no error during validation
	assert.NoError(t, err)

	// Check if the token contains the expected claim
	subClaim, ok := validatedToken.Claims.(jwt.MapClaims)["sub"].(string)
	assert.True(t, ok)
	assert.Equal(t, "sample", subClaim)
}

func TestTokenValidationFailsWithInvalidToken(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	invalidToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJpbnZhbGlkIiwiZXhwIjoxNjM0NTc1NTg3fQ.invalid_signature"

	_, err := tokenUtility.ValidateToken(invalidToken)
	assert.Error(t, err)
}

func TestTokenValidationFailsWithExpiredToken(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	claims := jwt.MapClaims{
		"sub": "expired",
		"exp": time.Now().Add(-1 * time.Minute).Unix(),
	}

	expiredToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedExpiredToken, _ := expiredToken.SignedString(tokenUtility.privateKey)

	_, err := tokenUtility.ValidateToken(signedExpiredToken)
	assert.Error(t, err)
}

func TestCanRefreshExpiredAccessToken(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	claims := map[string]interface{}{
		"sub": "sample",
	}

	accessToken, refreshToken, _ := tokenUtility.IssueToken(claims)

	// Make the access token expired
	expiredAccessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": "sample",
		"exp": time.Now().Add(-1 * time.Minute).Unix(),
	})
	signedExpiredAccessToken, _ := expiredAccessToken.SignedString(tokenUtility.GetPrivateKey())

	newAccessToken, newRefreshToken, err := tokenUtility.RefreshToken(signedExpiredAccessToken, refreshToken, time.Hour)

	// Check if there is no error during refreshing tokens
	assert.NoError(t, err)

	// Check if new access and refresh tokens are not empty and different from the old ones
	assert.NotEmpty(t, newAccessToken)
	assert.NotEmpty(t, newRefreshToken)
	assert.NotEqual(t, accessToken, newAccessToken)
	assert.NotEqual(t, refreshToken, newRefreshToken)
}
