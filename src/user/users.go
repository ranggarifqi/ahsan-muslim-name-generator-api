package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type IUserRepository interface {
	FindById(id string) (*User, error)
}

type IUserUsecase interface {
	FindById(id string) (*UserWithoutPassword, error)
	SignIn(payload *SignInDTO) (*SignInResult, error)
}
