package routes

import (
	"github.com/darth-raijin/bolig-side/api/controllers"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initializeProfile(api fiber.Router) {
	profile := api.Group("/profile")

	profile.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))
	profile.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	profile.Put("/", controllers.GetEventById)
	profile.Delete("/", controllers.GetEventById)
}
