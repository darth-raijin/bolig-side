package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Initialize() *fiber.App {
	app := fiber.New()

	if os.Getenv("ENV") != "prod" {
		initializeSwagger(app)
	}

	api := app.Group("/api")

	// Registering endpoints
	initializeAuth(api)
	initializeEvent(api)
	initializeFeedback(api)
	initializeProfile(api)

	return app
}
