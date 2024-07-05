package middleware

import (
	"bankApp1/internal/models"
	"context"
)

type UserRedisRepo interface {
	GetUserSession(ctx context.Context, sessionID string) (models.Claims, error)
}
