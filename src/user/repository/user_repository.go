package repository

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) user.IUserRepository {
	return &userRepository{
		conn,
	}
}

func (ur *userRepository) FindById(id string) (*user.User, error) {
	var res user.User
	err := ur.conn.First(&res, "id = ?", id).Error
	return &res, err
}
