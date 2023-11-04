package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	setupBucketRoutes(app)
	setupTranscoderRoutes(app)
	setupStreamRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Educado Transcoding API 👋!")
	})
}
