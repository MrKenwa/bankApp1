package usecase

import (
	"bankApp1/internal/balances/usecase"
	"bankApp1/internal/models"
	"context"
)

type (
	BalanceUC interface {
		Get(context.Context, *usecase.Filter) (models.Balance, error)
		Increase(context.Context, *usecase.Filter, int64) error
		Decrease(context.Context, *usecase.Filter, int64) error
	}

	OperationRepo interface {
		Create(context.Context, *models.Operation) (models.OperationID, error)
	}
)
