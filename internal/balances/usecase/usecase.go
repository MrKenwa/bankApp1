package usecase

import (
	"bankApp1/internal/models"
	"context"
	"fmt"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type BalanceUC struct {
	manager     *manager.Manager
	balanceRepo BalanceRepo
}

func NewBalanceUsecase(manager *manager.Manager, balanceRepo BalanceRepo) *BalanceUC {
	return &BalanceUC{
		manager:     manager,
		balanceRepo: balanceRepo,
	}
}

func (buc *BalanceUC) Create(ctx context.Context, b *CreateBalance) (models.BalanceID, error) {
	var bid models.BalanceID
	var err error
	if err := buc.manager.Do(ctx, func(ctx context.Context) error {
		balance := b.toBalance()
		bid, err = buc.balanceRepo.Create(ctx, &balance)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return bid, nil
}

func (buc *BalanceUC) Delete(ctx context.Context, b *Filter) error {
	if err := buc.manager.Do(ctx, func(ctx context.Context) error {
		filter := b.toBalanceFilter()
		err := buc.balanceRepo.Delete(ctx, &filter)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (buc *BalanceUC) Get(ctx context.Context, f *Filter) (models.Balance, error) {
	filter := f.toBalanceFilter()
	balance, err := buc.balanceRepo.Get(ctx, &filter)
	if err != nil {
		return models.Balance{}, err
	}
	return balance, nil
}

func (buc *BalanceUC) Increase(ctx context.Context, f *Filter, amount int64) error {
	fmt.Println("я в функции increase")
	if err := buc.manager.Do(ctx, func(ctx context.Context) error {
		filter := f.toBalanceFilter()
		err := buc.balanceRepo.Increase(ctx, &filter, amount)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (buc *BalanceUC) Decrease(ctx context.Context, f *Filter, amount int64) error {
	if err := buc.manager.Do(ctx, func(ctx context.Context) error {
		filter := f.toBalanceFilter()
		err := buc.balanceRepo.Decrease(ctx, &filter, amount)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
