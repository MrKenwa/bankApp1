package delivery

import "github.com/gofiber/fiber/v2"

func MapPaymentRoutes(group fiber.Router, h Handlers) {
	group.Post("/sendMoney", h.Send())
	group.Post("/payIn", h.PayIn())
	group.Post("/payOut", h.PayOut())
}
