package usecase

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
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
		DeletedAt: u.DeletedAt.Time,
	}

	return res, nil
}

func (uuc *userUsecase) SignIn(payload *user.SignInDTO) (*user.SignInResult, error) {
	foundUser, err := uuc.ur.FindOne("email = ?", payload.Email)
	if err != nil {
		return nil, err
	}

	_, err = helper.ComparePassword(foundUser.Password, payload.Password)
	if err != nil {
		return nil, err
	}

	res := &user.SignInResult{
		UserWithoutPassword: user.UserWithoutPassword{
			ID:        foundUser.ID,
			Email:     foundUser.Email,
			FullName:  foundUser.FullName,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
			DeletedAt: foundUser.DeletedAt.Time,
		},
	}

	return res, nil
}
