package productsDelivery

import (
	"bankApp1/internal/models"
	"bankApp1/internal/products/productsUsecase"
	"errors"
)

type CreateCardRequest struct {
	CardType string `json:"card_type"`
	Pin      string `json:"pin"`
}

type DeleteCardRequest struct {
	CardID models.CardID `json:"card_id"`
}

type CreateDepositRequest struct {
	Type    string  `json:"deposit_type"`
	IntRate float32 `json:"int_rate"`
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

func (req *CreateCardRequest) toCreateCard() productsUsecase.CreateCard {
	return productsUsecase.CreateCard{
		Type: req.CardType,
		Pin:  req.Pin,
	}
}

func (req *CreateDepositRequest) toCreateDeposit() productsUsecase.CreateDeposit {
	return productsUsecase.CreateDeposit{
		Type:    req.Type,
		IntRate: req.IntRate,
	}
}

func (req *CreateCardRequest) checkData() error {
	if req.CardType == "" || req.Pin == "" {
		return errors.New("invalid data")
	}
	return nil
}

func (req *CreateDepositRequest) checkData() error {
	if req.Type == "" || req.IntRate <= 1 {
		return errors.New("invalid data")
	}
	return nil
}
