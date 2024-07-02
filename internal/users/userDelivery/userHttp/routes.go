package userHttp

import (
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h Handlers) {
	group.Post("/register", h.Register())
	group.Get("/login", h.Login())
}
