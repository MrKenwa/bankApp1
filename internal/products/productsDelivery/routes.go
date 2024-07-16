package productsDelivery

import (
	"bankApp1/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapProductsRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/createCard", mw.AuthedMiddleware(), h.CreateNewCard())
	group.Post("/createDeposit", mw.AuthedMiddleware(), h.CreateNewDeposit())
	group.Delete("/deleteCard", mw.AuthedMiddleware(), h.DeleteCard())
	group.Delete("/deleteDeposit", mw.AuthedMiddleware(), h.DeleteDeposit())
	group.Get("/getCards", mw.AuthedMiddleware(), h.GetCards())
	group.Get("/getDeposits", mw.AuthedMiddleware(), h.GetDeposits())
}
