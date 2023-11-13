package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/api/v1/common"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func ListBucket(c *fiber.Ctx) error {

	// List files from GCP
	var list, err = gcp.Service.ListFiles()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.Error{
				Code:    "E0001",
				Message: err.Error(),
			},
		})
	}

	// Return list of files as JSON
	return c.Status(fiber.StatusOK).JSON(list)
}
