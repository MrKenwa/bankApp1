package productsUsecase

import (
	"bankApp1/internal/balances/balancesUsecase"
	"bankApp1/internal/models"
	"context"
	"errors"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type ProductsUC struct {
	manager   *manager.Manager
	cardUC    CardUC
	depositUC DepositUC
	balanceUC BalanceUC
}

func NewProductsUC(manager *manager.Manager, cardUC CardUC, depositUC DepositUC, balanceUC BalanceUC) ProductsUC {
	return ProductsUC{
		manager:   manager,
		cardUC:    cardUC,
		depositUC: depositUC,
		balanceUC: balanceUC,
	}
}

func (u *ProductsUC) CreateNewCard(ctx context.Context, cardData CreateCard) (models.CardID, error) {
	var cid models.CardID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		card := cardData.toCard()

		cid, err = u.cardUC.Create(ctx, card)
		if err != nil {
			return err
		}

		balance := balancesUsecase.CreateBalance{
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

func (u *ProductsUC) CreateNewDeposit(ctx context.Context, depositData CreateDeposit) (models.DepositID, error) {
	var did models.DepositID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposit := depositData.toDeposit()

		did, err = u.depositUC.Create(ctx, deposit)
		if err != nil {
			return err
		}

		balance := balancesUsecase.CreateBalance{
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

func (u *ProductsUC) GetManyCards(ctx context.Context, uid models.UserID) (models.ManyCards, error) {
	var cards models.ManyCards
	var err error
	filter := models.CardFilter{UserIDs: []models.UserID{uid}}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		cards, err = u.cardUC.GetMany(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return cards, nil
}

func (u *ProductsUC) GetManyDeposits(ctx context.Context, uid models.UserID) (models.ManyDeposits, error) {
	var deposits models.ManyDeposits
	var err error
	filter := models.DepositFilter{UserIDs: []models.UserID{uid}}
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposits, err = u.depositUC.GetMany(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return deposits, nil
}

func (u *ProductsUC) DeleteCard(ctx context.Context, cid models.CardID, uid models.UserID) error {
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		ok, err := u.isItUserCard(ctx, cid, uid)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("not enough rights")
		}

		err = u.cardUC.Delete(ctx, cid)
		if err != nil {
			return err
		}

		filter := models.BalanceFilter{CardIDs: []models.CardID{cid}}
		err = u.balanceUC.Delete(ctx, filter)
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

func (u *ProductsUC) DeleteDeposit(ctx context.Context, did models.DepositID, uid models.UserID) error {
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		ok, err := u.isItUserDeposit(ctx, did, uid)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("not enough rights")
		}

		err = u.depositUC.Delete(ctx, did)
		if err != nil {
			return err
		}

		filter := models.BalanceFilter{DepositIDs: []models.DepositID{did}}
		err = u.balanceUC.Delete(ctx, filter)
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

func (u *ProductsUC) isItUserCard(ctx context.Context, cid models.CardID, uid models.UserID) (bool, error) {
	card, err := u.cardUC.Get(ctx, models.CardFilter{IDs: []models.CardID{cid}})
	if err != nil {
		return false, err
	}

	if card.UserID == uid {
		return true, nil
	}
	return false, nil
}

func (u *ProductsUC) isItUserDeposit(ctx context.Context, did models.DepositID, uid models.UserID) (bool, error) {
	card, err := u.depositUC.Get(ctx, models.DepositFilter{IDs: []models.DepositID{did}})
	if err != nil {
		return false, err
	}

	if card.UserID == uid {
		return true, nil
	}
	return false, nil
}
