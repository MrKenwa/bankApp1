package cardsUsecase

import (
	"bankApp1/internal/models"
)

type CreateCard struct {
	Number int
	UserID models.UserID
	Type   string
	Pin    string
}

func (c *CreateCard) toCard() models.Card {
	return models.Card{
		CardNumber: c.Number,
		UserID:     c.UserID,
		Type:       c.Type,
		Pin:        c.Pin,
	}
}
