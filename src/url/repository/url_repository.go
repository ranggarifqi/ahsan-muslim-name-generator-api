package repository

import (
	"github.com/ranggarifqi/qashir-api/src/url"
	"gorm.io/gorm"
)

type urlRepository struct {
	conn *gorm.DB
}

func NewUrlRepository(conn *gorm.DB) url.IUrlRepository {
	return &urlRepository{
		conn,
	}
}

func (r *urlRepository) FindById(id string) (*url.Url, error) {
	var res url.Url
	err := r.conn.First(&res, "id=?", id).Error
	return &res, err
}

func (r *urlRepository) Create(payload *url.CreateUrlDto) (*url.Url, error) {
	user := url.Url{
		Url: payload.Url,
	}
	err := r.conn.Create(&user).Error
	return &user, err
}
