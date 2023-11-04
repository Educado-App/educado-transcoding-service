package stream

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Stream(c *fiber.Ctx) error {
	// Get fileName from URL
	fileName := c.Params("fileName")

	fmt.Printf("%v", fileName)

	return c.SendString("Streaming...")
}
