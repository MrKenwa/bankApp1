package models

import "time"

type Deposit struct {
	DepositID    DepositID  `db:"deposit_id"`
	UserID       UserID     `db:"user_id"`
	Type         string     `db:"deposit_type"`
	InterestRate float32    `db:"interest_rate"`
	CreatedAt    time.Time  `db:"created_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}

type ManyDeposits []Deposit

type DepositFilter struct {
	IDs           []DepositID
	UserIDs       []UserID
	Types         []string
	InterestRates []float32
}
