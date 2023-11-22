package transcoder

import (
	"bufio"
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
