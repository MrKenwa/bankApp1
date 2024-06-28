package userHttp

import (
	"bankApp1/internal/models"
	"bankApp1/internal/users/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	UserUC interface {
		Register(context.Context, *usecase.RegisterUser) (models.UserID, error)
		Login(context.Context, *usecase.LoginUser) (models.UserID, error)
	}

	Handlers interface {
		Register() fiber.Handler
		Login() fiber.Handler
	}
)
