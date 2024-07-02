package balancesUsecase

import (
	"bankApp1/internal/models"
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type BalanceUC struct {
	manager     *manager.Manager
	balanceRepo BalanceRepo
}

func NewBalanceUC(manager *manager.Manager, balanceRepo BalanceRepo) BalanceUC {
	return BalanceUC{
		manager:     manager,
		balanceRepo: balanceRepo,
	}
}

func (u *BalanceUC) Create(ctx context.Context, b CreateBalance) (models.BalanceID, error) {
	var bid models.BalanceID
	var err error

	balance := b.toBalance()
	bid, err = u.balanceRepo.Create(ctx, balance)
	if err != nil {
		return -1, err
	}

	return bid, nil
}

func (u *BalanceUC) Get(ctx context.Context, filter models.BalanceFilter) (models.Balance, error) {
	balance, err := u.balanceRepo.Get(ctx, filter)
	if err != nil {
		return models.Balance{}, err
	}
	return balance, nil
}

func (u *BalanceUC) Increase(ctx context.Context, filter models.BalanceFilter, amount int64) error {
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.balanceRepo.Increase(ctx, filter, amount)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (u *BalanceUC) Decrease(ctx context.Context, filter models.BalanceFilter, amount int64) error {
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.balanceRepo.Decrease(ctx, filter, amount)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (u *BalanceUC) Delete(ctx context.Context, filter models.BalanceFilter) error {
	err := u.balanceRepo.Delete(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
