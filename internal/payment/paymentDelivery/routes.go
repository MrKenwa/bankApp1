package paymentDelivery

import (
	"bankApp1/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapPaymentRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/sendMoney", mw.AuthedMiddleware(), h.Send())
	group.Post("/payIn", mw.AuthedMiddleware(), h.PayIn())
	group.Post("/payOut", mw.AuthedMiddleware(), h.PayOut())
}
