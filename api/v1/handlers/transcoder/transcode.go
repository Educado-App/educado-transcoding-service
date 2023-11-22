package transcoder

import (
	GCPService "github.com/Educado-App/educado-transcoding-service/internals/gcp"
	Service "github.com/Educado-App/educado-transcoding-service/internals/transcoder"
	"github.com/gofiber/fiber/v2"
	"os"
	"sync"
)

func Transcode(c *fiber.Ctx) error {
	//Extract the file from formData
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0002",
				"message": "Upload request does not contain a file",
			},
		})
	}

	//Extract the filename from formData
	filename := c.FormValue("fileName")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0003",
				"message": "Upload request does not contain a fileName",
			},
		})
	}

	//Make sure the file is an MP4
	allowedTypes := []string{"video/mp4"}
	contentType := file.Header.Get("Content-Type")
	if !contains(allowedTypes, contentType) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0004",
				"message": "File type not allowed",
			},
		})
	}

	// Remove .mp4 if present
	if filename[len(filename)-4:] == ".mp4" {
		filename = filename[:len(filename)-4]
	}

	// Make sure the temp dir exists
	if _, err := os.Stat("./tmp"); os.IsNotExist(err) {
		os.Mkdir("./tmp", 0755)
	}

	//Save the file to the local filesystem
	err = c.SaveFile(file, "./tmp/"+filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0005",
				"message": "Unable to save the file",
			},
		})
	}

	//Wait group to wait for all transcodes to finish
	wg := sync.WaitGroup{}

	//Use ffmpeg to transcode the file into 4 different resolutions: 1080p, 720p, 480p, 360p (reversed resolution dimensions)
	resolutions := []string{"360x640"}
	for _, resolution := range resolutions {
		outputPath := "./tmp/" + filename + "_" + resolution + ".mp4"
		wg.Add(1)
		go Service.TranscodeVideo("./tmp/"+filename, outputPath, resolution, &wg)
	}

	//Wait for all transcodes to finish
	wg.Wait()

	//Delete the original file
	err = os.Remove("./tmp/" + filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "E0006",
				"message": "Unable to delete the original file",
			},
		})
	}

	//Upload the transcoded files to the bucket
	for _, resolution := range resolutions {
		localFilePath := "./tmp/" + filename + "_" + resolution + ".mp4"
		file, err := Service.FileToBytes(localFilePath)
		err = GCPService.Service.UploadFile(filename+"_"+resolution+".mp4", file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    "E0007",
					"message": "Unable to upload the transcoded file",
				},
			})
		}
		//Delete the transcoded file
		err = os.Remove(localFilePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    "E0008",
					"message": "Unable to delete the transcoded file",
				},
			})
		}
	}

	//Return filename
	return c.JSON(fiber.Map{
		"filename": filename,
		"file":     file,
	})
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
