package main

import (
	database "customer-service/databases"
	"customer-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()

	app := fiber.New()

	routes.RouteInit(app)

	app.Listen(":5002")
}
