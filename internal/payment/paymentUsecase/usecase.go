package paymentUsecase

import (
	"bankApp1/internal/models"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type PaymentUC struct {
	manager   *manager.Manager
	opRepo    OperationRepo
	balanceUC BalanceUC
	cardUC    CardsUC
	depositUC DepositsUC
}

func NewPaymentUC(manager *manager.Manager, balanceUC BalanceUC, oRepo OperationRepo, cardUC CardsUC, depUC DepositsUC) *PaymentUC {
	return &PaymentUC{
		manager:   manager,
		balanceUC: balanceUC,
		opRepo:    oRepo,
		cardUC:    cardUC,
		depositUC: depUC,
	}
}

func (u *PaymentUC) Send(ctx context.Context, sendData *SendData) (models.OperationID, error) {
	var opid models.OperationID
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		senderFilter, receiverFilter := sendData.toBalanceFilter()
		senderBalance, err := u.balanceUC.Get(ctx, senderFilter)
		if err != nil {
			return err
		}

		ok, err := u.isItUserBalance(ctx, senderBalance.BalanceID, sendData.UserID)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("not enough rights")
		}

		if senderBalance.Amount-sendData.Amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balanceUC.Decrease(ctx, senderFilter, sendData.Amount); err != nil {
			return err
		}

		if err := u.balanceUC.Increase(ctx, receiverFilter, sendData.Amount); err != nil {
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

	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		balanceFilter := payData.toBalanceFilter()
		if err := u.balanceUC.Increase(ctx, balanceFilter, payData.Amount); err != nil {
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

	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		balanceFilter := payData.toBalanceFilter()
		senderBalance, err := u.balanceUC.Get(ctx, balanceFilter)
		if err != nil {
			return err
		}

		ok, err := u.isItUserBalance(ctx, senderBalance.BalanceID, payData.UserID)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("not enough rights")
		}

		if senderBalance.Amount-payData.Amount < 0 {
			return errors.New("not enough money")
		}

		if err := u.balanceUC.Decrease(ctx, balanceFilter, payData.Amount); err != nil {
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

func (u *PaymentUC) isItUserBalance(ctx context.Context, bid models.BalanceID, uid models.UserID) (bool, error) {
	balance, err := u.balanceUC.Get(ctx, models.BalanceFilter{IDs: []models.BalanceID{bid}})
	if err != nil {
		return false, err
	}

	if balance.CardID != nil {
		cid := *balance.CardID
		card, err := u.cardUC.Get(ctx, models.CardFilter{IDs: []models.CardID{cid}})
		if err != nil {
			return false, err
		}

		if card.UserID == uid {
			return true, nil
		}

		return false, nil
	}

	if balance.DepositID != nil {
		did := *balance.DepositID
		deposit, err := u.depositUC.Get(ctx, models.DepositFilter{IDs: []models.DepositID{did}})
		if err != nil {
			return false, err
		}

		if deposit.UserID == uid {
			return true, nil
		}

		return false, nil
	}

	return false, errors.New("balance not found")
}
