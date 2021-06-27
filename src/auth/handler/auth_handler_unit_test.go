package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
	authService "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth/services"
	authUC "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth/usecase"
	phService "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/passwordhasher/services"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	userRepo "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	myValidator "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/validator"
)

type Mocks struct {
	mockUserRepository *userRepo.UserRepositoryMock
	mockBcryptService  *phService.BCryptServiceMock
	mockJWTService     *authService.JwtServiceMock
}

func setupMock() Mocks {
	return Mocks{
		&userRepo.UserRepositoryMock{},
		&phService.BCryptServiceMock{},
		&authService.JwtServiceMock{},
	}
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Validator = myValidator.NewMyValidator()
	return e
}

func Test_Auth_Handler_SignIn(t *testing.T) {
	email := "test@ranggarifqi.com"
	id, _ := uuid.NewV4()
	fullName := "Test User"

	findOneResult := &user.User{
		ID:       id,
		Email:    email,
		Password: "thisisanencryptedpassword",
		FullName: fullName,
	}

	t.Run("Should return correct response on successful sign in", func(t *testing.T) {
		tokenResult := "thisisjwttoken"
		claim := &auth.AuthClaim{
			UserID: findOneResult.ID,
			Email:  findOneResult.Email,
		}

		/* Setup Mocks */
		mocks := setupMock()
		mocks.mockUserRepository.On("FindOne", "email = ?", []interface{}{email}).Return(findOneResult, nil)
		mocks.mockBcryptService.On("ComparePassword", findOneResult.Password, "correctpassword").Return(true, nil)
		mocks.mockJWTService.On("GetAuthToken", claim).Return(&tokenResult, nil)

		/* Setup Handler */
		usecase := authUC.NewAuthUsecase(mocks.mockUserRepository, mocks.mockJWTService, mocks.mockBcryptService)
		handler := &AuthHandler{
			authUsecase: usecase,
		}

		/* Setup request */
		e := setupEcho()
		jsonBody := `{"email": "test@ranggarifqi.com", "password": "correctpassword"}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/signin", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		err := handler.SignIn(ctx)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
