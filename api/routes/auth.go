package routes

import (
	"github.com/darth-raijin/bolig-side/api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initializeAuth(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Post("/login", controllers.RegisterUser)
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] auth log}\n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Get("/register", controllers.GetRegisterView)
	auth.Post("/register", controllers.RegisterUser)
}
