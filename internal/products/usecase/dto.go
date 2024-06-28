package usecase

import "bankApp1/internal/models"

type NewCard struct {
	UserID models.UserID
	Type   string
	Pin    string
}

func (c *NewCard) toCard() *models.Card {
	return &models.Card{
		UserID: c.UserID,
		Type:   c.Type,
		Pin:    c.Pin,
	}
}
