package models

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserID    UserID    `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`

	jwt.StandardClaims
}
