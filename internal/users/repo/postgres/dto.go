package postgres

import (
	"bankApp1/internal/models"
	"time"
)

type user struct {
	UserID         models.UserID `db:"id"`
	Name           string        `db:"name"`
	Lastname       string        `db:"lastname"`
	Patronymic     string        `db:"patronymic"`
	Email          string        `db:"email"`
	Password       string        `db:"password"`
	PassportNumber string        `db:"passport_number"`
	CreatedAt      time.Time     `db:"created_at"`
	DeletedAt      *time.Time    `db:"deleted_at"`
}

type manyUsers []user

func (u *user) toUser() models.User {
	return models.User{
		UserID:         u.UserID,
		Name:           u.Name,
		Lastname:       u.Lastname,
		Patronymic:     u.Patronymic,
		Email:          u.Email,
		Password:       u.Password,
		PassportNumber: u.PassportNumber,
		CreatedAt:      u.CreatedAt,
		DeletedAt:      u.DeletedAt,
	}
}

func (u manyUsers) toManyUsers() []models.User {
	users := make([]models.User, len(u))
	for _, user := range u {
		users = append(users, user.toUser())
	}
	return users
}
