package main

import (
	"fmt"
	"github.com/Educado-App/educado-transcoding-service/api/v1/routes"
	"github.com/Educado-App/educado-transcoding-service/config"
	"github.com/Educado-App/educado-transcoding-service/internals/gcp"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

var (
	Version = "dev"
	Build   = "none"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "transcoder-service",
		ServerHeader: "Edudcado/transcoder-service/" + Version + " (Build " + Build + ")",
	})

	files, err := gcp.Service.ListFiles()
	if err != nil {
		return
	}
	fmt.Printf("%v", files)

	//TODO: Add DB Inits

	// Mount routes.
	routes.SetupRoutes(app)

	//TODO: Add tls config stuff
	// Listen for incoming traffic.
	log.Fatal(app.Listen(":" + config.APIPort))
}
