package user

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	FullName  string         `json:"fullName" gorm:"column:fullName"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index;column:deletedAt"`
}

type IUserRepository interface {
	FindById(id string) (*User, error)
	FindOne(query interface{}, data ...interface{}) (*User, error)
}

type IUserUsecase interface {
	FindById(id string) (*UserWithoutPassword, error)
	SignIn(payload *SignInDTO) (*SignInResult, error)
}
