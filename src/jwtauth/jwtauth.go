package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type JwtClaims struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	jwt.StandardClaims
}
