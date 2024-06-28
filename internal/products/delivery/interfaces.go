package delivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/products/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductUC interface {
		CreateNewCard(context.Context, *usecase.CreateCard) (models.CardID, error)
		DeleteCard(context.Context, models.CardID) error
		CreateNewDeposit(context.Context, *usecase.CreateDeposit) (models.DepositID, error)
		DeleteDeposit(context.Context, models.DepositID) error
		GetCards(context.Context, models.UserID) (models.ManyCards, error)
		GetDeposits(context.Context, models.UserID) (models.ManyDeposits, error)
	}

	Handlers interface {
		CreateNewCard() fiber.Handler
		DeleteCard() fiber.Handler
		CreateNewDeposit() fiber.Handler
		DeleteDeposit() fiber.Handler
		GetCards() fiber.Handler
		GetDeposits() fiber.Handler
	}
)
