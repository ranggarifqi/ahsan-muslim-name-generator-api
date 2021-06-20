package url

import (
	"time"

	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"gorm.io/gorm"
)

type Url struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = helper.GenerateID(6)
	return nil
}

type CreateUrlDto struct {
	Url string `json:"url" form:"url" validate:"required,url"`
}

type IUrlRepository interface {
	FindById(id string) (*Url, error)
	Create(payload *CreateUrlDto) (*Url, error)
}

type IUrlUsecase interface {
	FindById(id string) (*Url, error)
	Create(payload *CreateUrlDto) (*Url, error)
}
