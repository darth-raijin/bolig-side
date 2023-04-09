package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initialize() *fiber.App {
	app := fiber.New()

	if os.Getenv("ENV") != "prod" {
		initializeSwagger(app)
	}

	api := app.Group("/api")

	api.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	// Registering endpoints
	initializeAuth(api)
	initializeProfile(api)

	return app
}
