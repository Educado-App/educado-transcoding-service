package bucket

import (
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func ListBucket(c *fiber.Ctx) error {
	var list, err = gcp.Service.ListFiles()
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(list)
}
