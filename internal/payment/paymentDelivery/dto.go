package paymentDelivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/payment/paymentUsecase"
	"errors"
)

type SendRequest struct {
	SenderBalanceID   models.BalanceID `json:"sender_balance_id"`
	ReceiverBalanceID models.BalanceID `json:"receiver_balance_id"`
	Amount            int64            `json:"amount"`
}

func (r SendRequest) checkData() error {
	if r.Amount <= 0 {
		return errors.New("invalid data")
	}
	if r.SenderBalanceID <= 0 {
		return errors.New("invalid data")
	}
	if r.ReceiverBalanceID <= 0 {
		return errors.New("invalid data")
	}
	return nil
}

func (r SendRequest) toSendData() *paymentUsecase.SendData {
	return &paymentUsecase.SendData{
		SendBalanceID:    &r.SenderBalanceID,
		ReceiveBalanceID: &r.ReceiverBalanceID,
		Amount:           r.Amount,
		OpType:           "transfer",
	}
}

type PayRequest struct {
	BalanceID models.BalanceID `json:"balance_id"`
	Amount    int64            `json:"amount"`
}

func (r PayRequest) toPayData() *paymentUsecase.PayData {
	return &paymentUsecase.PayData{
		BalanceID: r.BalanceID,
		Amount:    r.Amount,
	}
}

func (r PayRequest) checkData() error {
	if r.Amount <= 0 {
		return errors.New("invalid data")
	}
	if r.BalanceID <= 0 {
		return errors.New("invalid data")
	}
	return nil
}
