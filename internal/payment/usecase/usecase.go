package usecase

import (
	"bankApp1/internal/balances/usecase"
	"bankApp1/internal/models"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type PaymentUC struct {
	manager   *manager.Manager
	opRepo    OperationRepo
	balanceUC BalanceUC
}

func NewPaymentUC(manager *manager.Manager, balanceUC BalanceUC, oRepo OperationRepo) *PaymentUC {
	return &PaymentUC{
		manager:   manager,
		balanceUC: balanceUC,
		opRepo:    oRepo,
	}
}

func (u *PaymentUC) Send(ctx context.Context, sendData *SendData) (models.OperationID, error) {
	var opid models.OperationID
	if err := sendData.CheckData(); err != nil {
		return -1, err
	}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		senderFilter, receiverFilter := sendData.toBalanceFilter()
		senderBalance, err := u.balanceUC.Get(ctx, (*usecase.Filter)(senderFilter))
		if err != nil {
			return err
		}

		if senderBalance.Amount-sendData.Amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balanceUC.Decrease(ctx, (*usecase.Filter)(senderFilter), sendData.Amount); err != nil {
			return err
		}

		if err := u.balanceUC.Increase(ctx, (*usecase.Filter)(receiverFilter), sendData.Amount); err != nil {
			return err
		}

		operation := sendData.toOperation()
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

func (u *PaymentUC) PayIn(ctx context.Context, payData *PayData) (models.OperationID, error) {
	var opid models.OperationID
	var err error
	if err := payData.CheckData(); err != nil {
		return -1, err
	}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		balanceFilter := payData.toBalanceFilter()
		if err := u.balanceUC.Increase(ctx, (*usecase.Filter)(balanceFilter), payData.Amount); err != nil {
			return err
		}

		operation := payData.toOperation()

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

func (u *PaymentUC) PayOut(ctx context.Context, payData *PayData) (models.OperationID, error) {
	var opid models.OperationID
	if err := payData.CheckData(); err != nil {
		return -1, err
	}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		balanceFilter := payData.toBalanceFilter()
		senderBalance, err := u.balanceUC.Get(ctx, (*usecase.Filter)(balanceFilter))
		if err != nil {
			return err
		}

		if senderBalance.Amount-payData.Amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balanceUC.Decrease(ctx, (*usecase.Filter)(balanceFilter), payData.Amount); err != nil {
			return err
		}

		operation := payData.toOperation()

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
