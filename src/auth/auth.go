package auth

import "github.com/gofrs/uuid"

type AuthClaim struct {
	UserID uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
}

type IAuthService interface {
	GetAuthToken(claim *AuthClaim) (*string, error)
}

type IAuthUsecase interface {
	SignIn(payload *SignInDTO) (*SignInResult, error)
}
