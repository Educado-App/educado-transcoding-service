package bucket

import "github.com/gofiber/fiber/v2"

func ListBucket(c *fiber.Ctx) error {
	return c.SendString("ListBucket")
}
