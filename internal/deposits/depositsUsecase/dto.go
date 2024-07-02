package depositsUsecase

import "bankApp1/internal/models"

type CreateDeposit struct {
	UserID       models.UserID
	Type         string
	InterestRate float32
}

func (d CreateDeposit) toDeposit() models.Deposit {
	return models.Deposit{
		UserID:       d.UserID,
		Type:         d.Type,
		InterestRate: d.InterestRate,
	}
}
