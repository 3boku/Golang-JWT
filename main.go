package main

import (
	"github.com/3boku/Golang-JWT/database"
	"github.com/3boku/Golang-JWT/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
