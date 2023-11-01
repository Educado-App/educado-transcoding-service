package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func ListBucket(c *fiber.Ctx) error {

	// List files from GCP
	var list, err = gcp.Service.ListFiles()
	if err != nil {
		return c.SendString(err.Error())
	}

	// Return list of files as JSON
	return c.JSON(list)
}
