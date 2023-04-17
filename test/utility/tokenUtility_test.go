package test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/darth-raijin/bolig-side/middleware"
	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/gofiber/fiber/v2"
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
	signedExpiredToken, _ := expiredToken.SignedString(tokenUtility.GetPrivateKey())

	_, err := tokenUtility.ValidateToken(signedExpiredToken)
	assert.Error(t, err)
}

func TestCanRefreshExpiredAccessToken(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	claims := map[string]interface{}{
		"sub": "sample",
	}

	accessToken, refreshToken, _ := tokenUtility.IssueToken(claims)

	newAccessToken, newRefreshToken, err := tokenUtility.RefreshToken(refreshToken)

	// Check if there is no error during refreshing tokens
	assert.NoError(t, err)

	// Check if new access and refresh tokens are not empty and different from the old ones
	assert.NotEmpty(t, newAccessToken)
	assert.NotEmpty(t, newRefreshToken)
	assert.NotEqual(t, accessToken, newAccessToken)
}

func TestJwtValidationMiddleware(t *testing.T) {
	tokenUtility, _ := utility.GetTokenUtilityInstance()

	t.Run("missing authorization header", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.JwtValidationMiddleware(tokenUtility))
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		req := httptest.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(body), "Missing Authorization header")
	})

	t.Run("invalid authorization header format", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.JwtValidationMiddleware(tokenUtility))
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "InvalidAuthHeader")
		resp, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(body), "Invalid Authorization header format")
	})

	t.Run("invalid or expired token", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.JwtValidationMiddleware(tokenUtility))
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer InvalidToken")
		resp, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Contains(t, string(body), "Invalid or expired token")
	})

	t.Run("valid token", func(t *testing.T) {
		app := fiber.New()
		app.Use(middleware.JwtValidationMiddleware(tokenUtility))
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		validToken, _, _ := tokenUtility.IssueToken(jwt.MapClaims{})
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+validToken)
		resp, err := app.Test(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, "Hello, World!", strings.TrimSpace(string(body)))
	})
}
