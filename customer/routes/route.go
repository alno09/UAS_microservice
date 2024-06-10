package routes

import (
	"customer-service/handlers"
	"customer-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user/:id", middleware.Auth, handlers.GetUserById)
	r.Post("/user/register", handlers.Register)
	r.Post("/user/login", handlers.Login)
}
