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

	OperationRepo interface {
		Create(context.Context, models.Operation) (models.OperationID, error)
	}
)
