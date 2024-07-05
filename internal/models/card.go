package models

type Card struct {
	CardID     CardID  `db:"card_id"`
	CardNumber int     `db:"card_number"`
	UserID     UserID  `db:"user_id"`
	Type       string  `db:"card_type"`
	Pin        string  `db:"pin_code"`
	CreatedAt  string  `db:"created_at"`
	DeletedAt  *string `db:"deleted_at"`
}

type ManyCards []Card

type CardFilter struct {
	IDs     []CardID
	Numbers []CardNumber
	UserIDs []UserID
	Types   []string
}
