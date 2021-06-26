package service

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/passwordhasher"
	"golang.org/x/crypto/bcrypt"
)

type bcryptService struct{}

func NewBcryptService() passwordhasher.IPasswordHasherService {
	return &bcryptService{}
}

func (s *bcryptService) HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(result), err
}

func (s *bcryptService) ComparePassword(hashedPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
