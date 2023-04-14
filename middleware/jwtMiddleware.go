package middleware

import (
	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func jwtValidationMiddleware(tokenUtility *utility.TokenUtility) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header.
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token"})
		}

		// Validate the token.
		token, err := tokenUtility.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
		}

		// Extract the claims and add them to the fiber.Ctx.Locals map.
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid claims"})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
