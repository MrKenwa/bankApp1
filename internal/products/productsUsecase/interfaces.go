package productsUsecase

import (
	"bankApp1/internal/balances/balancesUsecase"
	"bankApp1/internal/cards/cardsUsecase"
	"bankApp1/internal/deposits/depositsUsecase"
	"bankApp1/internal/models"
	"context"
)

type (
	CardUC interface {
		Create(context.Context, cardsUsecase.CreateCard) (models.CardID, error)
		Get(context.Context, models.CardFilter) (models.Card, error)
		GetMany(context.Context, models.CardFilter) ([]models.Card, error)
		Delete(context.Context, models.CardID) error
	}

	DepositUC interface {
		Create(context.Context, depositsUsecase.CreateDeposit) (models.DepositID, error)
		Get(context.Context, models.DepositFilter) (models.Deposit, error)
		GetMany(context.Context, models.DepositFilter) ([]models.Deposit, error)
		Delete(context.Context, models.DepositID) error
	}

	BalanceUC interface {
		Create(context.Context, balancesUsecase.CreateBalance) (models.BalanceID, error)
		Delete(context.Context, models.BalanceFilter) error
	}
)
