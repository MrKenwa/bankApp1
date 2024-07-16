package redis

import (
	"bankApp1/config"
	"bankApp1/internal/models"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserRedisRepo struct {
	cfg *config.Config
	db  *redis.Client
}

func NewUserRedisRepo(cfg *config.Config, db *redis.Client) *UserRedisRepo {
	return &UserRedisRepo{cfg: cfg, db: db}
}

func (r *UserRedisRepo) SetUserSession(ctx context.Context, sessionID string, claims models.Claims) error {
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		return err
	}

	if err := r.db.Set(
		ctx,
		sessionID,
		claimsJson,
		time.Second*time.Duration(r.cfg.SessionSettings.SessionTTLSeconds),
	).Err(); err != nil {
		return err
	}
	return nil
}

func (r *UserRedisRepo) GetUserSession(ctx context.Context, sessionID string) (models.Claims, error) {
	var claimsJson []byte

	if err := r.db.Get(ctx, sessionID).Scan(&claimsJson); err != nil {
		return models.Claims{}, err
	}

	var claims models.Claims
	if err := json.Unmarshal(claimsJson, &claims); err != nil {
		return models.Claims{}, err
	}

	return claims, nil
}
