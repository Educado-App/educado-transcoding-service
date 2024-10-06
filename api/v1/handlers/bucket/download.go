package bucket

import (
	"encoding/base64"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/Educado-App/educado-transcoding-service/api/v1/common"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
)

func DownloadFile(c *fiber.Ctx) error {
    // Get filename from URL
    var filename = c.Params("fileName")

    // Download file from GCP
    var file, err = gcp.Service.DownloadFile(filename)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": common.Error{
                Code:    "E0001",
                Message: err.Error(),
            },
        })
    }

	// Get file attributes to set additional headers
    attrs, err := gcp.Service.GetFileAttributes(filename)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": common.Error{
                Code:    "E0002",
                Message: err.Error(),
            },
        })
    }

    // Determine the MIME type based on the file extension
    mimeType := mime.TypeByExtension(filepath.Ext(filename))
    if mimeType == "" {
		// If the file has no file extension, return the file as a base64 encoded string (Old way)
		// TODO: Remove this in the future, once all files have a file extension
		return c.SendString(base64.StdEncoding.EncodeToString(file))
    }

    // Set the Content-Type header
    c.Set(fiber.HeaderContentType, mimeType)

	// Set the Last-Modified header
    lastModified := attrs.Updated.Format(http.TimeFormat)
    c.Set(fiber.HeaderLastModified, lastModified)

    // Send the file as a binary response
    return c.Send(file)
}