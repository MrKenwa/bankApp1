package usecase

import (
	"bankApp1/internal/models"
	"context"
)

type (
	CardRepo interface {
		Create(ctx context.Context, c *models.Card) (models.CardID, error)
		GetMany(ctx context.Context, filter *models.CardFilter) (models.ManyCards, error)
		Delete(ctx context.Context, id models.CardID) error
	}

	DepositRepo interface {
		Create(ctx context.Context, d *models.Deposit) (models.DepositID, error)
		GetMany(ctx context.Context, filter *models.DepositFilter) (models.ManyDeposits, error)
		Delete(ctx context.Context, id models.DepositID) error
	}

	BalanceUC interface {
		Create(context.Context, *models.Balance) (models.BalanceID, error)
		Delete(context.Context, *models.BalanceFilter) error
	}
)
