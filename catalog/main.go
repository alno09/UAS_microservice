package main

import (
	"UAS-micro/config"
	"UAS-micro/database"
	"UAS-micro/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	config.RabbitMQInit()

	app := fiber.New()

	app.Get("/catalog", handler.GetAllCatalog)
	app.Get("/catalog/:id", handler.GetCatalogById)
	app.Post("/catalog", handler.CreateCatalog)

	app.Listen(":5001")
}
