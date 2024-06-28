package delivery

import "github.com/gofiber/fiber/v2"

func MapProductsRoutes(group fiber.Router, h Handlers) {
	group.Post("/createCard", h.CreateNewCard())
	group.Post("/createDeposit", h.CreateNewDeposit())
	group.Delete("/deleteCard", h.DeleteCard())
	group.Delete("/deleteDeposit", h.DeleteDeposit())
	group.Get("/getCards", h.GetCards())
	group.Get("/getDeposits", h.GetDeposits())
}
