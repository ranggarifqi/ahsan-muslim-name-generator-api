package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserWithoutPassword struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type SignInResult struct {
	UserWithoutPassword
	Token string `json:"token"`
}
