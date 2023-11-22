package transcoder

import (
	"log"
	"os/exec"
)

func TranscodeVideo(inputPath, outputPath, resolution string) {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-s", resolution, outputPath)
	err := cmd.Run()
	if err != nil {
		log.Printf("Error transcoder file: %s to %s, error: %v", inputPath, resolution, err)
		return
	}
	log.Printf("Transcoding to resolution %s completed: %s", resolution, outputPath)
}
