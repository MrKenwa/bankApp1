package balancesUsecase

import (
	"bankApp1/internal/models"
)

type CreateBalance struct {
	CardID    *models.CardID
	DepositID *models.DepositID
	Amount    int64
}

func (b *CreateBalance) toBalance() models.Balance {
	return models.Balance{
		CardID:    b.CardID,
		DepositID: b.DepositID,
		Amount:    b.Amount,
	}
}

type Filter struct {
	IDs        []models.BalanceID
	CardIDs    []models.CardID
	DepositIDs []models.DepositID
}

func (b *Filter) toBalanceFilter() models.BalanceFilter {
	return models.BalanceFilter{
		IDs:        b.IDs,
		CardIDs:    b.CardIDs,
		DepositIDs: b.DepositIDs,
	}
}
