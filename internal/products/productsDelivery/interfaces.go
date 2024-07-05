package productsDelivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/products/productsUsecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductsUC interface {
		CreateNewCard(context.Context, productsUsecase.CreateCard) (models.CardID, error)
		CreateNewDeposit(context.Context, productsUsecase.CreateDeposit) (models.DepositID, error)
		GetManyCards(context.Context, models.UserID) (models.ManyCards, error)
		GetManyDeposits(context.Context, models.UserID) (models.ManyDeposits, error)
		DeleteCard(context.Context, models.CardID, models.UserID) error
		DeleteDeposit(context.Context, models.DepositID, models.UserID) error
	}

	Handlers interface {
		CreateNewCard() fiber.Handler
		CreateNewDeposit() fiber.Handler
		GetCards() fiber.Handler
		GetDeposits() fiber.Handler
		DeleteCard() fiber.Handler
		DeleteDeposit() fiber.Handler
	}
)
