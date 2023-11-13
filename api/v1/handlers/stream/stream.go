package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	"io"
)

func Stream(c *fiber.Ctx) error {
	fileName := c.Params("fileName")

	// Get the file attributes to check the Content-Type
	attrs, err := gcp.Service.Attributes(fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0001",
				"message": err.Error(),
			},
		})
	}

	// Check if the Content-Type is "video/mp4"
	if attrs.ContentType != "video/mp4" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0008",
				"message": "File is not an MP4 video",
			},
		})
	}

	// Create reader for file
	reader, err := gcp.Service.Reader(fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0001",
				"message": err.Error(),
			},
		})
	}
	defer reader.Close()

	// Set the Content-Type header for the response
	c.Set("Content-Type", attrs.ContentType)

	// Stream file to client
	if _, err = io.Copy(c.Response().BodyWriter(), reader); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0001",
				"message": err.Error(),
			},
		})
	}

	return nil
}
