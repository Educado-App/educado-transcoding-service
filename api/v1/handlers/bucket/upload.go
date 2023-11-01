package bucket

import (
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {

	return c.SendString("UploadFile")
}
