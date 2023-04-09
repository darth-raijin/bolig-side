package routes

import (
	"github.com/gofiber/fiber/v2"
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
}
