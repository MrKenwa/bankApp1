package delivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/products/usecase"
)

type CreateCardRequest struct {
	UserID   models.UserID `json:"user_id"`
	CardType string        `json:"card_type"`
	Pin      string        `json:"pin"`
}

func (req *CreateCardRequest) toCreateCard() *usecase.CreateCard {
	return &usecase.CreateCard{
		UserID: req.UserID,
		Type:   req.CardType,
		Pin:    req.Pin,
	}
}

type DeleteCardRequest struct {
	CardID models.CardID `json:"card_id"`
}

type CreateDepositRequest struct {
	UserID  models.UserID `json:"user_id"`
	Type    string        `json:"type"`
	IntRate float32       `json:"int_rate"`
}

func (req *CreateDepositRequest) toCreateDeposit() *usecase.CreateDeposit {
	return &usecase.CreateDeposit{
		UserID:  req.UserID,
		Type:    req.Type,
		IntRate: req.IntRate,
	}
}

type DeleteDepositRequest struct {
	DepositID models.DepositID `json:"deposit_id"`
}

type GetCardsRequest struct {
	UserID models.UserID `json:"user_id"`
}

type GetDepositsRequest struct {
	UserID models.UserID `json:"user_id"`
}
