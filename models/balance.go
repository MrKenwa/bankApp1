package models

import "time"

type Balance struct {
	BalanceID BalanceID  `db:"balance_id"`
	CardID    *CardID    `db:"card_id"`
	DepositID *DepositID `db:"deposit_id"`
	Amount    int64      `db:"amount"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type ManyBalances []Balance

type BalanceFilter struct {
	IDs        []BalanceID
	CardIDs    []CardID
	DepositIDs []DepositID
}
