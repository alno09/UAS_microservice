package route

import (
	"order-service/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/orders", handler.CreateOrder)
	r.Get("/orders", handler.FindOrders)
}
