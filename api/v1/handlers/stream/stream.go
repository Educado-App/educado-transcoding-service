package stream

import (
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	"io"
)

func Stream(c *fiber.Ctx) error {
	// Get fileName from URL
	fileName := c.Params("fileName")

	// Create reader for file
	var reader, err = gcp.Service.Reader(fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer reader.Close()

	// Set the Content-Type header
	c.Set("Content-Type", "video/mp4")

	// Stream file to client
	if _, err := io.Copy(c, reader); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return nil
}
