package main

import (
	"log"
	"registration-app/database"
	"registration-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	routes.SetupAuth(app)
	routes.SetupUser(app)

	log.Fatal(app.Listen(":3000"))
}
