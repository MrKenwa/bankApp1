package productsUsecase

import (
	"bankApp1/models"
	"bankApp1/txManager"
	"context"
	"time"
)

type CardRepo interface {
	Create(ctx context.Context, c models.Card) (models.CardID, error)
	GetMany(ctx context.Context, filter models.CardFilter) (models.ManyCards, error)
	Delete(ctx context.Context, id models.CardID) error
}

type DepositRepo interface {
	Create(ctx context.Context, d models.Deposit) (models.DepositID, error)
	GetMany(ctx context.Context, filter models.DepositFilter) (models.ManyDeposits, error)
	Delete(ctx context.Context, id models.DepositID) error
}

type BalanceRepo interface {
	Create(ctx context.Context, b models.Balance) (models.BalanceID, error)
	Delete(ctx context.Context, filter models.BalanceFilter) error
}

type ProductsUsecase struct {
	manager     *txManager.TxManager
	cardRepo    CardRepo
	depositRepo DepositRepo
	balanceRepo BalanceRepo
}

func NewProductsUsecase(manager *txManager.TxManager, cRepo CardRepo, dRepo DepositRepo, bRepo BalanceRepo) *ProductsUsecase {
	return &ProductsUsecase{
		manager:     manager,
		cardRepo:    cRepo,
		depositRepo: dRepo,
		balanceRepo: bRepo,
	}
}

func (u *ProductsUsecase) CreateNewCard(uid models.UserID, cType string, pin string) (models.CardID, error) {
	ctx := context.Background()
	var cid models.CardID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		card := models.Card{
			UserID: uid,
			Type:   cType,
			Pin:    pin,
		}

		cid, err = u.cardRepo.Create(ctx, card)
		if err != nil {
			return err
		}

		balance := models.Balance{
			CardID:    &cid,
			Amount:    0,
			CreatedAt: time.Now(),
		}
		_, err := u.balanceRepo.Create(ctx, balance)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return cid, nil
}

func (u *ProductsUsecase) DeleteCard(cid models.CardID) error {
	ctx := context.Background()
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.cardRepo.Delete(ctx, cid)
		if err != nil {
			return err
		}

		filter := models.BalanceFilter{CardIDs: []models.CardID{cid}}
		err = u.balanceRepo.Delete(ctx, filter)
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

func (u *ProductsUsecase) CreateNewDeposit(uid models.UserID, dType string, intRate float32) (models.DepositID, error) {
	ctx := context.Background()
	var did models.DepositID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposit := models.Deposit{
			UserID:       uid,
			Type:         dType,
			InterestRate: intRate,
		}

		did, err = u.depositRepo.Create(ctx, deposit)
		if err != nil {
			return err
		}

		balance := models.Balance{
			DepositID: &did,
			Amount:    0,
			CreatedAt: time.Now(),
		}
		_, err := u.balanceRepo.Create(ctx, balance)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return did, nil
}

func (u *ProductsUsecase) DeleteDeposit(did models.DepositID) error {
	ctx := context.Background()
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

func (u *ProductsUsecase) GetCards(uid models.UserID) (models.ManyCards, error) {
	ctx := context.Background()
	var cards models.ManyCards
	var err error
	filter := models.CardFilter{UserIDs: []models.UserID{uid}}
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

func (u *ProductsUsecase) GetDeposits(uid models.UserID) (models.ManyDeposits, error) {
	ctx := context.Background()
	var deposits models.ManyDeposits
	var err error
	filter := models.DepositFilter{UserIDs: []models.UserID{uid}}
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
