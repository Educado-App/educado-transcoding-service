package stream

import "github.com/gofiber/fiber/v2"

func Stream(c *fiber.Ctx) error {
	return c.SendString("Streaming...")
}
