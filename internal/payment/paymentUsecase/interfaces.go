package paymentUsecase

import (
	"bankApp1/internal/models"
	"context"
)

type (
	BalanceUC interface {
		Get(context.Context, models.BalanceFilter) (models.Balance, error)
		Increase(context.Context, models.BalanceFilter, int64) error
		Decrease(context.Context, models.BalanceFilter, int64) error
	}

	CardsUC interface {
		Get(context.Context, models.CardFilter) (models.Card, error)
	}

	DepositsUC interface {
		Get(context.Context, models.DepositFilter) (models.Deposit, error)
	}

	OperationRepo interface {
		Create(context.Context, models.Operation) (models.OperationID, error)
	}
)
