package models

import "time"

type User struct {
	UserID         UserID     `db:"id"`
	Name           string     `db:"name"`
	Lastname       string     `db:"lastname"`
	Patronymic     string     `db:"patronymic"`
	Email          string     `db:"email"`
	Password       string     `db:"password"`
	PassportNumber string     `db:"passport_number"`
	CreatedAt      time.Time  `db:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at"`
}

type ManyUsers []User

type UserFilter struct {
	IDs         []UserID
	Names       []string
	Lastnames   []string
	Patronymics []string
	Emails      []string
}
