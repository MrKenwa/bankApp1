package balancesUsecase

import (
	"bankApp1/internal/models"
	"context"
)

type BalanceRepo interface {
	Create(context.Context, models.Balance) (models.BalanceID, error)
	Delete(context.Context, models.BalanceFilter) error
	Get(context.Context, models.BalanceFilter) (models.Balance, error)
	Decrease(context.Context, models.BalanceFilter, int64) error
	Increase(context.Context, models.BalanceFilter, int64) error
}
