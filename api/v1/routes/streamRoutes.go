package routes

import (
	handlers "github.com/Educado-App/educado-transcoding-service/api/v1/handlers/stream"
	"github.com/gofiber/fiber/v2"
)

func setupStreamRoutes(app *fiber.App) {
	bucket := app.Group("/stream")
	bucket.Get("/:fileName", handlers.Stream)
}
