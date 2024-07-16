package userUsecase

import (
	"bankApp1/internal/models"
	"context"
)

type (
	UserRepo interface {
		Create(ctx context.Context, u models.User) (models.UserID, error)
		Get(ctx context.Context, filter models.UserFilter) (models.User, error)
	}
	UserRedisRepo interface {
		SetUserSession(context.Context, string, models.Claims) error
		GetUserSession(context.Context, string) (models.Claims, error)
	}
)
