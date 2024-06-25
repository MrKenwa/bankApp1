package models

import "time"

type Operation struct {
	OperationID       OperationID `db:"operation_id"`
	SenderBalanceID   *BalanceID  `db:"sender_balance_id"`
	ReceiverBalanceID *BalanceID  `db:"receiver_balance_id"`
	Amount            int64       `db:"amount"`
	OperationType     string      `db:"operation_type"`
	CreatedAt         time.Time   `db:"created_at"`
	DeletedAt         *time.Time  `db:"deleted_at"`
}

type ManyOperations []Operation

type OperationFilter struct {
	IDs               []OperationID
	SenderBalanceID   []BalanceID
	ReceiverBalanceID []BalanceID
	OperationType     []string
}
