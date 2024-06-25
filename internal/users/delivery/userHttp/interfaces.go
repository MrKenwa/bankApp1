package userHttp

import (
	"bankApp1/internal/models"
	"bankApp1/internal/users/usecase"
	"github.com/gofiber/fiber/v2"
)

type (
	UserUC interface {
		Register(user *usecase.RegisterUser) (models.UserID, error)
		//Login(username, password string) (models.UserID, error)
	}

	Handlers interface {
		Register() fiber.Handler
	}
)
