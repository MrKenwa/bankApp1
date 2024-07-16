package paymentUsecase

import (
	"bankApp1/internal/models"
	"bankApp1/pkg"
	"time"
)

type SendData struct {
	UserID           models.UserID
	SendBalanceID    *models.BalanceID
	ReceiveBalanceID *models.BalanceID
	Amount           int64
	OpType           string
}

func (s *SendData) toBalanceFilter() (models.BalanceFilter, models.BalanceFilter) {
	sender := models.BalanceFilter{
		IDs: []models.BalanceID{*s.SendBalanceID},
	}
	receiver := models.BalanceFilter{
		IDs: []models.BalanceID{*s.ReceiveBalanceID},
	}
	return sender, receiver
}

func (s *SendData) toOperation() models.Operation {
	return models.Operation{
		SenderBalanceID:   s.SendBalanceID,
		ReceiverBalanceID: s.ReceiveBalanceID,
		Amount:            s.Amount,
		OperationType:     s.OpType,
		CreatedAt:         time.Now(),
	}
}

type PayData struct {
	UserID    models.UserID
	BalanceID models.BalanceID
	Amount    int64
	OpType    string
}

func (p *PayData) toBalanceFilter() models.BalanceFilter {
	return models.BalanceFilter{
		IDs: []models.BalanceID{p.BalanceID},
	}
}

func (p *PayData) toOperation() models.Operation {
	var operation models.Operation
	if p.OpType == pkg.PayInOperationType {
		operation.ReceiverBalanceID = &p.BalanceID
	} else if p.OpType == pkg.PayOutOperationType {
		operation.SenderBalanceID = &p.BalanceID
	}
	operation.Amount = p.Amount
	operation.OperationType = p.OpType
	operation.CreatedAt = time.Now()
	return operation
}
