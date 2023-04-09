package routes

import (
	"github.com/darth-raijin/bolig-side/api/controllers"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initializeAuth(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	auth.Post("/login", controllers.RegisterUser)
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] auth log}\n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Get("/register", controllers.GetEventById)
	auth.Post("/register", controllers.RegisterUser)
}
