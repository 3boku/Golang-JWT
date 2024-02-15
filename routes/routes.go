package routes

import (
	"github.com/3boku/Golang-JWT/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Register)
}
