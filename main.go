package main

import (
	"go-api/configs"
	"go-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configs.ConnectDB()

	routes.UserRoute(app)
	app.Listen(":" + configs.ApiPort())
}
