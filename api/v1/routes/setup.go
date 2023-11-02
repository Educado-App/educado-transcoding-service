package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	setupBucketRoutes(app)
	setupTranscoderRoutes(app)
	setupStreamRoutes(app)
}
