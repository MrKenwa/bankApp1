package productsUsecase

import (
	"bankApp1/internal/cards/cardsUsecase"
	"bankApp1/internal/deposits/depositsUsecase"
	"bankApp1/internal/models"
)

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

func (c *CreateCard) toCard() cardsUsecase.CreateCard {
	return cardsUsecase.CreateCard{
		UserID: c.UserID,
		Type:   c.Type,
		Pin:    c.Pin,
	}
}

func (d *CreateDeposit) toDeposit() depositsUsecase.CreateDeposit {
	return depositsUsecase.CreateDeposit{
		UserID:       d.UserID,
		Type:         d.Type,
		InterestRate: d.IntRate,
	}
}
