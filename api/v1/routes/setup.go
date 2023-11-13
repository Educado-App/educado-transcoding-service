package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	setupBucketRoutes(api)
	setupTranscoderRoutes(api)
	setupStreamRoutes(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Educado Transcoding API 👋!")
	})
}
