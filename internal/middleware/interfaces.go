package middleware

import (
	"bankApp1/internal/models"
	"context"
)

type UserRepo interface {
	Get(context.Context, models.UserFilter) (models.User, error)
}
