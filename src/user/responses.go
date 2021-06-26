package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserWithoutPassword struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SignInResult struct {
	UserWithoutPassword
	Token string `json:"token"`
}
