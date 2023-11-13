package bucket

import (
	"fmt"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	"io"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0002",
				"message": "Upload request does not contain a file",
			},
		})
	}

	filename := c.FormValue("fileName")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0003",
				"message": "Upload request does not contain a fileName",
			},
		})
	}

	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png", "video/mp4"}
	contentType := file.Header.Get("Content-Type")
	if !contains(allowedTypes, contentType) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0004",
				"message": "File type not allowed",
			},
		})
	}

	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0005",
				"message": "Unable to open the file",
			},
		})
	}
	defer fileData.Close()

	content, err := io.ReadAll(fileData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0006",
				"message": "Unable to read the file",
			},
		})
	}

	err = gcp.Service.UploadFile(filename, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0007",
				"message": fmt.Sprintf("Failed to upload the file: %v", err),
			},
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("File %s uploaded successfully", filename),
	})
}

func contains(types []string, get string) bool {
	for _, t := range types {
		if t == get {
			return true
		}
	}
	return false
}
