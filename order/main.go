package main

import (
	"order-service/config"
	"order-service/database"
	"order-service/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()

	app := fiber.New()

	config.RabbitMQInit()

	route.RouteInit(app)

	app.Listen(":5003")
}
