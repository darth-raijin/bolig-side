package routes

import (
	"github.com/darth-raijin/bolig-side/api/controllers"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initializeEvent(api fiber.Router) {
	event := api.Group("/event")

	event.Use(logger.New(logger.Config{
		Format: "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}\n" +
			"${resBody}\n",
		TimeZone: "Europe/Copenhagen",
	}))

	event.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	event.Post("/", controllers.CreateEvent)
}
