package usecase

import (
	"errors"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
	authService "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth/services"
	phService "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/passwordhasher/services"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	userRepo "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	"github.com/stretchr/testify/assert"
)

func Test_Auth_Usecase_SignIn(t *testing.T) {
	t.Run("Should return correct value on successfull sign in", func(t *testing.T) {
		mockUserRepository := &userRepo.UserRepositoryMock{}
		mockBcryptService := &phService.BCryptServiceMock{}
		mockJWTService := &authService.JwtServiceMock{}

		email := "test@ranggarifqi.com"
		signInDTO := auth.SignInDTO{
			Email:    email,
			Password: "rawpassword",
		}

		id, _ := uuid.NewV4()
		fullName := "Test User"
		tokenResult := "thisisjwttoken"

		findOneResult := &user.User{
			ID:       id,
			Email:    email,
			Password: "thisisanencryptedpassword",
			FullName: fullName,
		}

		expectedResult := &auth.SignInResult{
			UserWithoutPassword: user.UserWithoutPassword{
				ID:       id,
				Email:    email,
				FullName: fullName,
			},
			Token: tokenResult,
		}

		mockUserRepository.On("FindOne", "email = ?", []interface{}{email}).Return(findOneResult, nil)
		mockBcryptService.On("ComparePassword", findOneResult.Password, signInDTO.Password).Return(true, nil)

		claim := &auth.AuthClaim{
			UserID: findOneResult.ID,
			Email:  findOneResult.Email,
		}

		mockJWTService.On("GetAuthToken", claim).Return(&tokenResult, nil)

		usecase := NewAuthUsecase(mockUserRepository, mockJWTService, mockBcryptService)

		result, err := usecase.SignIn(&signInDTO)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Should return error if wrong email", func(t *testing.T) {
		mockUserRepository := &userRepo.UserRepositoryMock{}
		mockBcryptService := &phService.BCryptServiceMock{}
		mockJWTService := &authService.JwtServiceMock{}

		email := "random.email@asd.com"

		signInDTO := auth.SignInDTO{
			Email:    email,
			Password: "rawpassword",
		}

		errMsg := "record not found"
		findOneError := errors.New(errMsg)

		mockUserRepository.On("FindOne", "email = ?", []interface{}{email}).Return(nil, findOneError)

		usecase := NewAuthUsecase(mockUserRepository, mockJWTService, mockBcryptService)

		result, err := usecase.SignIn(&signInDTO)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errMsg)
	})
}
