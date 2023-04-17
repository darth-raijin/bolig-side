package routes

import (
	"os"

	"github.com/darth-raijin/bolig-side/middleware"
	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initialize(tokenUtility *utility.TokenUtility) *fiber.App {
	app := fiber.New()

	if os.Getenv("ENV") != "prod" {
		initializeSwagger(app)
	}

	api := app.Group("/api", middleware.JwtValidationMiddleware(tokenUtility))

	api.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	// Registering endpoints
	initializeAuth(api)
	initializeProfile(api)

	return app
}
