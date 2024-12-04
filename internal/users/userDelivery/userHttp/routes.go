package userHttp

import (
	"bankApp1/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/register", h.Register())
	group.Post("/login", h.Login())
	group.Post("/getOwn", mw.AuthedMiddleware(), h.GetOwn())
	group.Post("/hello", h.Hello())
}
