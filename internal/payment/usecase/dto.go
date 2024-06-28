package usecase

import (
	"bankApp1/internal/models"
	"errors"
	"time"
)

type SendData struct {
	SendBalanceID    *models.BalanceID
	ReceiveBalanceID *models.BalanceID
	Amount           int64
	OpType           string
}

func (s *SendData) toBalanceFilter() (*models.BalanceFilter, *models.BalanceFilter) {
	sender := &models.BalanceFilter{
		IDs: []models.BalanceID{*s.SendBalanceID},
	}
	receiver := &models.BalanceFilter{
		IDs: []models.BalanceID{*s.ReceiveBalanceID},
	}
	return sender, receiver
}

func (s *SendData) toOperation() *models.Operation {
	return &models.Operation{
		SenderBalanceID:   s.SendBalanceID,
		ReceiverBalanceID: s.ReceiveBalanceID,
		Amount:            s.Amount,
		OperationType:     s.OpType,
		CreatedAt:         time.Now(),
	}
}

type PayData struct {
	BalanceID *models.BalanceID
	Amount    int64
	OpType    string
}

func (p *PayData) toBalanceFilter() *models.BalanceFilter {
	return &models.BalanceFilter{
		IDs: []models.BalanceID{*p.BalanceID},
	}
}

func (p *PayData) toOperation() *models.Operation {
	var operation models.Operation
	if p.OpType == models.PayInOperationType {
		operation.ReceiverBalanceID = p.BalanceID
	} else if p.OpType == models.PayOutOperationType {
		operation.SenderBalanceID = p.BalanceID
	}
	operation.Amount = p.Amount
	operation.OperationType = p.OpType
	operation.CreatedAt = time.Now()
	return &operation
}

func (p *PayData) CheckData() error {
	if p.Amount <= 0 {
		return errors.New("invalid data")
	}
	if *p.BalanceID <= 0 {
		return errors.New("invalid data")
	}
	return nil
}

func (s *SendData) CheckData() error {
	if s.Amount <= 0 {
		return errors.New("invalid data")
	}
	if *s.SendBalanceID <= 0 {
		return errors.New("invalid data")
	}
	if *s.ReceiveBalanceID <= 0 {
		return errors.New("invalid data")
	}
	return nil
}
