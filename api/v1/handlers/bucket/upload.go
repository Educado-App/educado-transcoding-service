package bucket

import (
	"fmt"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	"io"
)

func UploadFile(c *fiber.Ctx) error {
	// Get the file from the POST request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Upload request does not contain a file")
	}

	// Get the filename from the formdata "filename" key
	filename := c.FormValue("fileName")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Upload request does not contain a fileName")
	}

	// Check if file is of the allowed types
	imageTypes := []string{"image/jpeg", "image/jpg", "image/png"}
	videoTypes := []string{"video/mp4"}
	isVideo := false
	if !contains(imageTypes, file.Header.Get("Content-Type")) {
		if !contains(videoTypes, file.Header.Get("Content-Type")) {
			return c.Status(fiber.StatusBadRequest).SendString("File type not allowed")
		}
		isVideo = true
	}

	if isVideo {
		println("Video file")
	} else {
		println("Image file")
	}

	// Open the file
	fileData, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to open the file")
	}
	defer fileData.Close()

	// Read the content of the file
	content, err := io.ReadAll(fileData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to read the file")
	}

	// Upload the file using GCPService
	err = gcp.Service.UploadFile(filename, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to upload the file: %v", err))
	}

	// Return success message
	return c.SendString(fmt.Sprintf("File %s uploaded successfully", filename))
}

// Check if a string is in a slice of strings
func contains(types []string, get string) bool {
	for _, t := range types {
		if t == get {
			return true
		}
	}
	return false
}
