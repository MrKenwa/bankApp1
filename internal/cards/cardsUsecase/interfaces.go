package cardsUsecase

import (
	"bankApp1/internal/models"
	"context"
)

type CardRepo interface {
	Create(context.Context, models.Card) (models.CardID, error)
	Get(context.Context, models.CardFilter) (models.Card, error)
	GetMany(context.Context, models.CardFilter) (models.ManyCards, error)
	Delete(context.Context, models.CardID) error
}
