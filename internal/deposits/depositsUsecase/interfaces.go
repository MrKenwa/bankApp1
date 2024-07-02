package depositsUsecase

import (
	"bankApp1/internal/models"
	"context"
)

type DepositsRepo interface {
	Create(context.Context, models.Deposit) (models.DepositID, error)
	Get(context.Context, models.DepositFilter) (models.Deposit, error)
	GetMany(context.Context, models.DepositFilter) (models.ManyDeposits, error)
	Delete(context.Context, models.DepositID) error
}
