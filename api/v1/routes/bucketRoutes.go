package routes

import (
	handlers "github.com/Educado-App/educado-transcoding-service/api/v1/handlers/bucket"
	"github.com/gofiber/fiber/v2"
)

func setupBucketRoutes(app *fiber.App) {
	bucket := app.Group("/bucket")
	bucket.Get("/", handlers.ListBucket)
}
