package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
)

type jwtService struct{}

type JwtClaim struct {
	auth.AuthClaim
	jwt.StandardClaims
}

func NewJwtService() auth.IAuthService {
	return &jwtService{}
}

func (s *jwtService) GetAuthToken(claim *auth.AuthClaim) (*string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	jwtClaim := &JwtClaim{
		AuthClaim: auth.AuthClaim{
			UserID: claim.UserID,
			Email:  claim.Email,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	result, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &result, nil
}
