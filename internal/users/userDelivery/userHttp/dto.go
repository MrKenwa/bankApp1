package userHttp

import (
	"bankApp1/internal/models"
	"bankApp1/internal/users/userUsecase"
)

type RegisterRequest struct {
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	Patronymic     string `json:"patronymic"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PassportNumber string `json:"passport_number"`
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

func (r *RegisterRequest) checkData() bool {
	if r.Name == "" || r.LastName == "" || r.Patronymic == "" || r.PassportNumber == "" || r.Email == "" || r.Password == "" {
		return false
	}
	return true
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

type getUserResponse struct {
	UserID     models.UserID
	Name       string
	Lastname   string
	Patronymic string
	Email      string
}
