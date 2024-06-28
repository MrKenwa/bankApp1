package usecase

import "bankApp1/internal/models"

type CreateCard struct {
	UserID models.UserID
	Type   string
	Pin    string
}

type CreateDeposit struct {
	UserID  models.UserID
	Type    string
	IntRate float32
}

func (c *CreateCard) toCard() *models.Card {
	return &models.Card{
		UserID: c.UserID,
		Type:   c.Type,
		Pin:    c.Pin,
	}
}

func (d *CreateDeposit) toDeposit() *models.Deposit {
	return &models.Deposit{
		UserID:       d.UserID,
		Type:         d.Type,
		InterestRate: d.IntRate,
	}
}
