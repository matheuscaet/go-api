package routes

import (
	"go-api/controllers"
	"go-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", middlewares.Auth, controllers.CreateUser)
	app.Get("/user/:userId", middlewares.Auth, controllers.GetAUser)
	app.Put("/user/:userId", middlewares.Auth, controllers.EditAUser)
	app.Delete("/user/:userId", middlewares.Auth, controllers.DeleteAUser)
	app.Get("/users", middlewares.Auth, controllers.GetAllUsers)
}
