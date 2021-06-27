package usecase

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
)

type userUsecase struct {
	ur user.IUserRepository
}

func NewUserUsecase(ur user.IUserRepository) user.IUserUsecase {
	return &userUsecase{
		ur,
	}
}

func (uuc *userUsecase) FindById(id string) (*user.UserWithoutPassword, error) {
	u, err := uuc.ur.FindById(id)
	if err != nil {
		return nil, err
	}

	res := &user.UserWithoutPassword{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	return res, nil
}
