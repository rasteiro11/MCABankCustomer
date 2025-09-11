package domain

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
