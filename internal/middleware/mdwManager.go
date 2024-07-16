package middleware

import (
	"bankApp1/config"
	"bankApp1/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type MDWManager struct {
	cfg      *config.Config
	userRepo UserRepo
}

func NewMDWManager(cfg *config.Config, userRepo UserRepo) *MDWManager {
	return &MDWManager{cfg: cfg, userRepo: userRepo}
}

func (m *MDWManager) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		if headers == nil || len(headers) == 0 {
			return errors.New("header is nil")
		}
		authHeader := headers["Authorization"]
		if authHeader == nil || len(authHeader) == 0 {
			return errors.New("header is nil")
		}

		header := authHeader[0]

		if header == "" {
			return errors.New("authorization header is empty")
		}

		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return errors.New("invalid authorization header")
		}

		token := parts[1]
		claims := &models.Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return m.cfg.PublicKey, nil
		})
		if err != nil {
			return err
		}

		if claims.ExpiresAt.Before(time.Now()) && c.Path() != "/users/refresh" {
			return errors.New("token is expired")
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
