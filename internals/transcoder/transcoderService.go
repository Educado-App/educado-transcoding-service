package transcoder

import (
	"bufio"
	GCPService "github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/exec"
	"sync"
)

func TranscodeVideo(inputPath, outputPath, resolution string, wg *sync.WaitGroup) {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-s", resolution, outputPath)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error transcoder file: %s to %s, error: %v", inputPath, resolution, err)
		return
	}
	log.Printf("Transcoding to resolution %s completed: %s", resolution, outputPath)
}

func FileToBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	return bytes, err
}

func TranscodeAndUpload(resolutions []string, filename string, c *fiber.Ctx) error {
	wg := sync.WaitGroup{}

	//Use ffmpeg to transcode the file into 4 different resolutions: 1080p, 720p, 480p, 360p (reversed resolution dimensions)
	for _, resolution := range resolutions {
		outputPath := "./tmp/" + filename + "_" + resolution + ".mp4"
		wg.Add(1)
		go TranscodeVideo("./tmp/"+filename, outputPath, resolution, &wg)
	}

	//Wait for all transcodes to finish
	wg.Wait()

	//Delete the original file
	var err = os.Remove("./tmp/" + filename)
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
		file, err := FileToBytes(localFilePath)
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

	return nil
}
