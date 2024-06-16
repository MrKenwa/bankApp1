package paymentUsecase

import (
	"bankApp1/models"
	"bankApp1/txManager"
	"context"
	"errors"
	"time"
)

type BalanceRepo interface {
	Get(ctx context.Context, filter models.BalanceFilter) (models.Balance, error)
	Increase(ctx context.Context, filter models.BalanceFilter, amount int64) error
	Decrease(ctx context.Context, filter models.BalanceFilter, amount int64) error
}

type OperationRepo interface {
	Create(ctx context.Context, op models.Operation) (models.OperationID, error)
}

type PaymentUC struct {
	manager *txManager.TxManager
	balRepo BalanceRepo
	opRepo  OperationRepo
}

func NewPaymentUC(manager *txManager.TxManager, bRepo BalanceRepo, oRepo OperationRepo) *PaymentUC {
	return &PaymentUC{
		manager: manager,
		balRepo: bRepo,
		opRepo:  oRepo,
	}
}

func (u *PaymentUC) Send(sFilter models.BalanceFilter, rFilter models.BalanceFilter, amount int64, opType string) (models.OperationID, error) {
	ctx := context.Background()
	var opid models.OperationID
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		senderBalance, err := u.balRepo.Get(ctx, sFilter)
		if err != nil {
			return err
		}
		receiverBalance, err := u.balRepo.Get(ctx, rFilter)
		if err != nil {
			return err
		}

		if senderBalance.Amount-amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balRepo.Decrease(ctx, sFilter, amount); err != nil {
			return err
		}

		if err := u.balRepo.Increase(ctx, rFilter, amount); err != nil {
			return err
		}

		operation := models.Operation{
			SenderBalanceID:   &senderBalance.BalanceID,
			ReceiverBalanceID: &receiverBalance.BalanceID,
			Amount:            amount,
			OperationType:     opType,
			CreatedAt:         time.Now(),
		}
		opid, err = u.opRepo.Create(ctx, operation)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return -1, err
	}
	return opid, nil
}

func (u *PaymentUC) PayIn(rFilter models.BalanceFilter, amount int64, opType string) (models.OperationID, error) {
	ctx := context.Background()
	var opid models.OperationID
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		receiverBalance, err := u.balRepo.Get(ctx, rFilter)
		if err != nil {
			return err
		}

		if err := u.balRepo.Increase(ctx, rFilter, amount); err != nil {
			return err
		}

		operation := models.Operation{
			ReceiverBalanceID: &receiverBalance.BalanceID,
			Amount:            amount,
			OperationType:     opType,
			CreatedAt:         time.Now(),
		}
		opid, err = u.opRepo.Create(ctx, operation)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return -1, err
	}
	return opid, nil
}

func (u *PaymentUC) PayOut(sFilter models.BalanceFilter, amount int64, opType string) (models.OperationID, error) {
	ctx := context.Background()
	var opid models.OperationID
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		senderBalance, err := u.balRepo.Get(ctx, sFilter)
		if err != nil {
			return err
		}

		if senderBalance.Amount-amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balRepo.Decrease(ctx, sFilter, amount); err != nil {
			return err
		}

		operation := models.Operation{
			SenderBalanceID: &senderBalance.BalanceID,
			Amount:          amount,
			OperationType:   opType,
			CreatedAt:       time.Now(),
		}
		opid, err = u.opRepo.Create(ctx, operation)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return -1, err
	}
	return opid, nil
}
