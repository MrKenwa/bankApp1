package userHttp

import "bankApp1/internal/users/usecase"

type RegisterRequest struct {
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	Patronymic     string `json:"patronymic"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PassportNumber string `json:"passportNumber"`
}

func (r *RegisterRequest) toRegisterUser() *usecase.RegisterUser {
	return &usecase.RegisterUser{
		Name:           r.Name,
		LastName:       r.LastName,
		Patronymic:     r.Patronymic,
		Email:          r.Email,
		Password:       r.Password,
		PassportNumber: r.PassportNumber,
	}
}
