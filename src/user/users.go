package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type IUserRepository interface {
	FindById(id string) (*User, error)
	FindOne(query interface{}, data ...interface{}) (*User, error)
}

type IUserUsecase interface {
	FindById(id string) (*UserWithoutPassword, error)
	SignIn(payload *SignInDTO) (*SignInResult, error)
}
