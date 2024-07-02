package userHttp

import "bankApp1/internal/users/userUsecase"

type RegisterRequest struct {
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	Patronymic     string `json:"patronymic"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PassportNumber string `json:"passportNumber"`
}

func (r *RegisterRequest) toRegisterUser() userUsecase.RegisterUser {
	return userUsecase.RegisterUser{
		Name:           r.Name,
		LastName:       r.LastName,
		Patronymic:     r.Patronymic,
		Email:          r.Email,
		Password:       r.Password,
		PassportNumber: r.PassportNumber,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) toLoginUser() userUsecase.LoginUser {
	return userUsecase.LoginUser{
		Email:    r.Email,
		Password: r.Password,
	}
}
