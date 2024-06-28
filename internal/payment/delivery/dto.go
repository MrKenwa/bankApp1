package delivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/payment/usecase"
)

type SendRequest struct {
	SenderBalanceID   models.BalanceID `json:"sender_balance_id"`
	ReceiverBalanceID models.BalanceID `json:"receiver_balance_id"`
	Amount            int64            `json:"amount"`
	OperationType     string           `json:"operation_type"`
}

func (r *SendRequest) toSendData() *usecase.SendData {
	return &usecase.SendData{
		SendBalanceID:    &r.SenderBalanceID,
		ReceiveBalanceID: &r.ReceiverBalanceID,
		Amount:           r.Amount,
		OpType:           r.OperationType,
	}
}

type PayRequest struct {
	BalanceID     *models.BalanceID `json:"balance_id"`
	Amount        int64             `json:"amount"`
	OperationType string            `json:"operation_type"`
}

func (p *PayRequest) toPayData() *usecase.PayData {
	return &usecase.PayData{
		BalanceID: p.BalanceID,
		Amount:    p.Amount,
		OpType:    p.OperationType,
	}
}
