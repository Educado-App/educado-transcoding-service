package bucket

import (
	"encoding/base64"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func DownloadFile(c *fiber.Ctx) error {
	// Get filename from URL
	var filename = c.Params("fileName")

	// Download file from GCP
	var file, err = gcp.Service.DownloadFile(filename)
	if err != nil {
		return c.SendString(err.Error())
	}

	// Encode file to base64 for frontend
	fileb64 := base64.StdEncoding.EncodeToString(file)
	return c.SendString(fileb64)
}
