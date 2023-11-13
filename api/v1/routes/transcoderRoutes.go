package routes

import (
	handlers "github.com/Educado-App/educado-transcoding-service/api/v1/handlers/transcoder"
	"github.com/gofiber/fiber/v2"
)

func setupTranscoderRoutes(app fiber.Router) {
	transcoder := app.Group("/transcoder")
	transcoder.Post("/", handlers.Transcode)
}
