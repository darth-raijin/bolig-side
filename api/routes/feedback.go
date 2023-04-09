package routes

import (
	"github.com/darth-raijin/bolig-side/api/controllers"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initializeFeedback(api fiber.Router) {
	feedback := api.Group("/feedback")

	feedback.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	feedback.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	feedback.Post("/", controllers.GetEventById)
	feedback.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] auth log}\n",
		TimeZone: "Europe/Copenhagen",
	}))
}
