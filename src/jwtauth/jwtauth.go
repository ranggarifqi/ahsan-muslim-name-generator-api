package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type JwtClaim struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}

type IJwtAuthService interface {
	GetJWTToken(claim *JwtClaim) (string, error)
}
