package routes

import (
	"os"

	"github.com/darth-raijin/bolig-side/pkg/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initialize(tokenUtility *utility.TokenUtility) *fiber.App {
	app := fiber.New()

	if os.Getenv("ENV") != "prod" {
		// Swagger is not protected by Auth middleware
		initializeSwagger(app)
	}

	//  middleware.JwtValidationMiddleware(tokenUtility)
	// Restructure to use middleware in appropiate place
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
