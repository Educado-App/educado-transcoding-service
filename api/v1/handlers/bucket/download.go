package bucket

import (
	"encoding/base64"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func DownloadFile(c *fiber.Ctx) error {
	var filename = c.Params("fileName")

	var file, err = gcp.Service.DownloadFile(filename)
	if err != nil {
		return c.SendString(err.Error())
	}

	fileb64 := base64.StdEncoding.EncodeToString(file)
	return c.SendString(fileb64)
}
