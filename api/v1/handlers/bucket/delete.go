package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/api/v1/common"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func DeleteFile(c *fiber.Ctx) error {
	// Get fileName from URL
	fileName := c.Params("fileName")

	// Delete file from GCP
	err := gcp.Service.DeleteFile(fileName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.Error{
				Code:    "E0001",
				Message: err.Error(),
			},
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File deleted: " + fileName,
	})
}
