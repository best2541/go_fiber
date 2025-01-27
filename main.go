package main

import (
	"go_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.Try)

	app.Get("/foods", controllers.GetAllFoods)

	app.Get("/foods/:id", controllers.GetFood)

	app.Post("/foods", controllers.InsertFood)

	app.Post("/foods/:id", controllers.EditFoods)

	app.Delete("/foods/:id", controllers.DeleteFood)

	app.Get("/info", controllers.GetInfo)

	app.Listen(":3000")
}
