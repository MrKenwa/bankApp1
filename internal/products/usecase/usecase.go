package usecase

import (
	"bankApp1/internal/balances/usecase"
	"bankApp1/internal/models"
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type ProductsUC struct {
	manager     *manager.Manager
	cardRepo    CardRepo
	depositRepo DepositRepo
	balanceUC   BalanceUC
}

func NewProductsUC(manager *manager.Manager, cRepo CardRepo, dRepo DepositRepo, balanceUC BalanceUC) *ProductsUC {
	return &ProductsUC{
		manager:     manager,
		cardRepo:    cRepo,
		depositRepo: dRepo,
		balanceUC:   balanceUC,
	}
}

func (u *ProductsUC) CreateNewCard(ctx context.Context, cardData *CreateCard) (models.CardID, error) {
	var cid models.CardID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		card := cardData.toCard()

		cid, err = u.cardRepo.Create(ctx, card)
		if err != nil {
			return err
		}

		balance := &usecase.CreateBalance{
			CardID: &cid,
			Amount: 0,
		}

		_, err := u.balanceUC.Create(ctx, balance)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return cid, nil
}

func (u *ProductsUC) DeleteCard(ctx context.Context, cid models.CardID) error {
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.cardRepo.Delete(ctx, cid)
		if err != nil {
			return err
		}

		filter := &models.BalanceFilter{CardIDs: []models.CardID{cid}}
		err = u.balanceUC.Delete(ctx, (*usecase.Filter)(filter))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductsUC) CreateNewDeposit(ctx context.Context, depositData *CreateDeposit) (models.DepositID, error) {
	var did models.DepositID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposit := depositData.toDeposit()

		did, err = u.depositRepo.Create(ctx, deposit)
		if err != nil {
			return err
		}

		balance := &usecase.CreateBalance{
			DepositID: &did,
			Amount:    0,
		}
		_, err := u.balanceUC.Create(ctx, balance)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return did, nil
}

func (u *ProductsUC) DeleteDeposit(ctx context.Context, did models.DepositID) error {
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.depositRepo.Delete(ctx, did)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductsUC) GetCards(ctx context.Context, uid models.UserID) (models.ManyCards, error) {
	var cards models.ManyCards
	var err error
	filter := &models.CardFilter{UserIDs: []models.UserID{uid}}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		cards, err = u.cardRepo.GetMany(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return cards, nil
}

func (u *ProductsUC) GetDeposits(ctx context.Context, uid models.UserID) (models.ManyDeposits, error) {
	var deposits models.ManyDeposits
	var err error
	filter := &models.DepositFilter{UserIDs: []models.UserID{uid}}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposits, err = u.depositRepo.GetMany(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return deposits, nil
}
