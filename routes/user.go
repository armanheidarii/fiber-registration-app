package routes

import (
	"registration-app/handlers"
	"registration-app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupUser(app *fiber.App) {
	app.Get("/get-all-profiles", middlewares.JWTMiddleware, handlers.GetUsers)
	app.Get("/get-profile/:username", middlewares.JWTMiddleware, handlers.GetUser)
	app.Put("/update-profile/:username", middlewares.JWTMiddleware, handlers.UpdateUser)
	app.Delete("/delete-account/:username", middlewares.JWTMiddleware, handlers.DeleteUser)
}
