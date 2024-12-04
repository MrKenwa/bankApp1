package userHttp

import (
	"bankApp1/internal/models"
	"bankApp1/internal/users/userUsecase"
	"context"
	"github.com/gofiber/fiber/v2"
)

type (
	UserUC interface {
		Register(context.Context, userUsecase.RegisterUser) (models.UserID, error)
		Login(context.Context, userUsecase.LoginUser) (string, error)
		GetUser(context.Context, models.UserID) (models.User, error)
	}

	Handlers interface {
		Register() fiber.Handler
		Login() fiber.Handler
		GetOwn() fiber.Handler
		Hello() fiber.Handler
	}
)
