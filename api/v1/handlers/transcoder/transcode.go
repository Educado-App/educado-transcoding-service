package transcoder

import "github.com/gofiber/fiber/v2"

func Transcode(c *fiber.Ctx) error {
	return c.SendString("Transcode")
}
