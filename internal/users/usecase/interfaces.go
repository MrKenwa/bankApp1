package usecase

import (
	"bankApp1/internal/models"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, u *models.User) (models.UserID, error)
	Get(ctx context.Context, filter *models.UserFilter) (models.User, error)
}
