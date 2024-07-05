package userHttp

import (
	"bankApp1/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/register", h.Register())
	group.Get("/login", h.Login())
	group.Get("/getOwn", mw.AuthedMiddleware(), h.GetOwn())
	group.Get("/refresh", mw.AuthedMiddleware(), h.RefreshToken())
}
