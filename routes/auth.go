package routes

import (
	"registration-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuth(app *fiber.App) {
	app.Post("/signup", handlers.CreateUser)
	app.Post("/login", handlers.Login)
}
