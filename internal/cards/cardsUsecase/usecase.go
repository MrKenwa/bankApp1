package cardsUsecase

import (
	"bankApp1/internal/models"
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type CardsUC struct {
	manager  *manager.Manager
	cardRepo CardRepo
}

func NewCardUC(manager *manager.Manager, cardRepo CardRepo) CardsUC {
	return CardsUC{
		manager:  manager,
		cardRepo: cardRepo,
	}
}

func (u *CardsUC) Create(ctx context.Context, c CreateCard) (models.CardID, error) {
	card := c.toCard()
	cid, err := u.cardRepo.Create(ctx, card)
	if err != nil {
		return -1, err
	}
	return cid, nil
}

func (u *CardsUC) Get(ctx context.Context, filter models.CardFilter) (models.Card, error) {
	card, err := u.cardRepo.Get(ctx, filter)
	if err != nil {
		return models.Card{}, err
	}
	return card, nil
}

func (u *CardsUC) GetMany(ctx context.Context, filter models.CardFilter) ([]models.Card, error) {
	var cards []models.Card
	var err error

	cards, err = u.cardRepo.GetMany(ctx, filter)
	if err != nil {
		return []models.Card{}, err
	}
	return cards, nil
}

func (u *CardsUC) Delete(ctx context.Context, cid models.CardID) error {
	err := u.cardRepo.Delete(ctx, cid)
	if err != nil {
		return err
	}
	return nil
}
