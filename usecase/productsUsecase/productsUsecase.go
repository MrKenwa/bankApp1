package productsUsecase

import (
	models2 "bankApp1/internal/models"
	"bankApp1/txManager"
	"context"
	"time"
)

type CardRepo interface {
	Create(ctx context.Context, c models2.Card) (models2.CardID, error)
	GetMany(ctx context.Context, filter models2.CardFilter) (models2.ManyCards, error)
	Delete(ctx context.Context, id models2.CardID) error
}

type DepositRepo interface {
	Create(ctx context.Context, d models2.Deposit) (models2.DepositID, error)
	GetMany(ctx context.Context, filter models2.DepositFilter) (models2.ManyDeposits, error)
	Delete(ctx context.Context, id models2.DepositID) error
}

type BalanceRepo interface {
	Create(ctx context.Context, b models2.Balance) (models2.BalanceID, error)
	Delete(ctx context.Context, filter models2.BalanceFilter) error
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

func (u *ProductsUsecase) CreateNewCard(uid models2.UserID, cType string, pin string) (models2.CardID, error) {
	ctx := context.Background()
	var cid models2.CardID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		card := models2.Card{
			UserID: uid,
			Type:   cType,
			Pin:    pin,
		}

		cid, err = u.cardRepo.Create(ctx, card)
		if err != nil {
			return err
		}

		balance := models2.Balance{
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

func (u *ProductsUsecase) DeleteCard(cid models2.CardID) error {
	ctx := context.Background()
	err := u.manager.Do(ctx, func(ctx context.Context) error {
		err := u.cardRepo.Delete(ctx, cid)
		if err != nil {
			return err
		}

		filter := models2.BalanceFilter{CardIDs: []models2.CardID{cid}}
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

func (u *ProductsUsecase) CreateNewDeposit(uid models2.UserID, dType string, intRate float32) (models2.DepositID, error) {
	ctx := context.Background()
	var did models2.DepositID
	var err error
	if err := u.manager.Do(ctx, func(ctx context.Context) error {
		deposit := models2.Deposit{
			UserID:       uid,
			Type:         dType,
			InterestRate: intRate,
		}

		did, err = u.depositRepo.Create(ctx, deposit)
		if err != nil {
			return err
		}

		balance := models2.Balance{
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

func (u *ProductsUsecase) DeleteDeposit(did models2.DepositID) error {
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

func (u *ProductsUsecase) GetCards(uid models2.UserID) (models2.ManyCards, error) {
	ctx := context.Background()
	var cards models2.ManyCards
	var err error
	filter := models2.CardFilter{UserIDs: []models2.UserID{uid}}
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

func (u *ProductsUsecase) GetDeposits(uid models2.UserID) (models2.ManyDeposits, error) {
	ctx := context.Background()
	var deposits models2.ManyDeposits
	var err error
	filter := models2.DepositFilter{UserIDs: []models2.UserID{uid}}
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
