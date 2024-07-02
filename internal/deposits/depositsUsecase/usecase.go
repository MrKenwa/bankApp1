package depositsUsecase

import (
	"bankApp1/internal/models"
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type DepositsUC struct {
	manager      *manager.Manager
	depositsRepo DepositsRepo
}

func NewDepositsUC(manager *manager.Manager, depositsRepo DepositsRepo) DepositsUC {
	return DepositsUC{
		manager:      manager,
		depositsRepo: depositsRepo,
	}
}

func (u *DepositsUC) Create(ctx context.Context, d CreateDeposit) (models.DepositID, error) {
	deposit := d.toDeposit()
	did, err := u.depositsRepo.Create(ctx, deposit)
	if err != nil {
		return -1, err
	}
	return did, nil
}

func (u *DepositsUC) Get(ctx context.Context, filter models.DepositFilter) (models.Deposit, error) {
	deposit, err := u.depositsRepo.Get(ctx, filter)
	if err != nil {
		return models.Deposit{}, err
	}
	return deposit, nil
}

func (u *DepositsUC) GetMany(ctx context.Context, filter models.DepositFilter) ([]models.Deposit, error) {
	deposits, err := u.depositsRepo.GetMany(ctx, filter)
	if err != nil {
		return []models.Deposit{}, err
	}
	return deposits, nil
}

func (u *DepositsUC) Delete(ctx context.Context, did models.DepositID) error {
	err := u.depositsRepo.Delete(ctx, did)
	if err != nil {
		return err
	}
	return nil
}
