package userHttp

import (
	"bankApp1/internal/models"
	"bankApp1/internal/users/usecase"
	"github.com/gofiber/fiber/v2"
)

type (
	UserUC interface {
		Register(user *usecase.RegisterUser) (models.UserID, error)
		Login(user *usecase.LoginUser) (models.UserID, error)
	}

	Handlers interface {
		Register() fiber.Handler
		Login() fiber.Handler
	}
)
