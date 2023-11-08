package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func DeleteFile(c *fiber.Ctx) error {
	// Get fileName from URL
	fileName := c.Params("fileName")

	// Delete file from GCP
	err := gcp.Service.DeleteFile(fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return success message
	return c.SendString("File deleted: " + fileName)
}
