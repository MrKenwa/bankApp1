package usecase

import "bankApp1/internal/models"

type RegisterUser struct {
	Name           string
	LastName       string
	Patronymic     string
	Email          string
	Password       string
	PassportNumber string
}

func (r *RegisterUser) toUser() *models.User {
	return &models.User{
		Name:           r.Name,
		Lastname:       r.LastName,
		Patronymic:     r.Patronymic,
		Email:          r.Email,
		Password:       r.Password,
		PassportNumber: r.PassportNumber,
	}
}
