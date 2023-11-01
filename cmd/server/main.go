package main

import (
	"github.com/Educado-App/educado-transcoding-service/api/v1/routes"
	"github.com/Educado-App/educado-transcoding-service/config"
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
		ServerHeader: "Educado/transcoder-service/" + Version + " (Build " + Build + ")",
	})

	//TODO: Add DB Inits

	// Mount routes.
	routes.SetupRoutes(app)

	//TODO: Add tls config stuff
	// Listen for incoming traffic.
	log.Fatal(app.Listen(":" + config.APIPort))
}
