package usecase

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/passwordhasher"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
)

type authUsecase struct {
	userRepository        user.IUserRepository
	authService           auth.IAuthService
	passwordHasherService passwordhasher.IPasswordHasherService
}

func NewAuthUsecase(ur user.IUserRepository, authService auth.IAuthService, phs passwordhasher.IPasswordHasherService) auth.IAuthUsecase {
	return &authUsecase{
		userRepository:        ur,
		authService:           authService,
		passwordHasherService: phs,
	}
}

func (uc *authUsecase) SignIn(payload *auth.SignInDTO) (*auth.SignInResult, error) {
	foundUser, err := uc.userRepository.FindOne("email = ?", payload.Email)
	if err != nil {
		return nil, err
	}

	_, err = uc.passwordHasherService.ComparePassword(foundUser.Password, payload.Password)
	if err != nil {
		return nil, err
	}

	claims := &auth.AuthClaim{
		UserID: foundUser.ID,
		Email:  foundUser.Email,
	}

	token, err := uc.authService.GetAuthToken(claims)
	if err != nil {
		return nil, err
	}

	res := &auth.SignInResult{
		UserWithoutPassword: user.UserWithoutPassword{
			ID:        foundUser.ID,
			Email:     foundUser.Email,
			FullName:  foundUser.FullName,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
		},
		Token: *token,
	}

	return res, nil
}
