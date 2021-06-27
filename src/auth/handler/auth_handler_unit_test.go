package handler

import (
	"encoding/json"
	"errors"
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
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/response"
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

		/* Assertion */
		if assert.NoError(t, handler.SignIn(ctx)) {
			res := response.SuccessResponse{}
			json.Unmarshal([]byte(rec.Body.String()), &res)
			resData, _ := res.Data.(map[string]interface{})

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Equal(t, "User signed in successfully", res.Message)
			assert.Equal(t, tokenResult, resData["token"])
		}
	})

	t.Run("Should return bad request error if password not specified", func(t *testing.T) {
		/* Setup Mocks */
		mocks := setupMock()

		/* Setup Handler */
		usecase := authUC.NewAuthUsecase(mocks.mockUserRepository, mocks.mockJWTService, mocks.mockBcryptService)
		handler := &AuthHandler{
			authUsecase: usecase,
		}

		/* Setup request */
		e := setupEcho()
		jsonBody := `{"email": "test@ranggarifqi.com"}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/signin", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		/* Assertion */
		if assert.NoError(t, handler.SignIn(ctx)) {
			res := response.ErrorResponse{}
			json.Unmarshal([]byte(rec.Body.String()), &res)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
			assert.Equal(t, "Key: 'SignInDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag", res.Message)
		}
	})

	t.Run("Should return bad request error if invalid email", func(t *testing.T) {
		/* Setup Mocks */
		mocks := setupMock()

		/* Setup Handler */
		usecase := authUC.NewAuthUsecase(mocks.mockUserRepository, mocks.mockJWTService, mocks.mockBcryptService)
		handler := &AuthHandler{
			authUsecase: usecase,
		}

		/* Setup request */
		e := setupEcho()
		jsonBody := `{"email": "notanemail", "password": "correctpassword"}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/signin", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		/* Assertion */
		if assert.NoError(t, handler.SignIn(ctx)) {
			res := response.ErrorResponse{}
			json.Unmarshal([]byte(rec.Body.String()), &res)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
			assert.Equal(t, "Key: 'SignInDTO.Email' Error:Field validation for 'Email' failed on the 'email' tag", res.Message)
		}
	})

	t.Run("Should return unauthorized error if using wrong credential", func(t *testing.T) {
		/* Setup Mocks */
		mocks := setupMock()
		mocks.mockUserRepository.On("FindOne", "email = ?", []interface{}{email}).Return(findOneResult, nil)
		mocks.mockBcryptService.On("ComparePassword", findOneResult.Password, "incorrectpassword").Return(false, errors.New("adsadsad"))

		/* Setup Handler */
		usecase := authUC.NewAuthUsecase(mocks.mockUserRepository, mocks.mockJWTService, mocks.mockBcryptService)
		handler := &AuthHandler{
			authUsecase: usecase,
		}

		/* Setup request */
		e := setupEcho()
		jsonBody := `{"email": "test@ranggarifqi.com", "password": "incorrectpassword"}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/signin", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		/* Assertion */
		if assert.NoError(t, handler.SignIn(ctx)) {
			res := response.ErrorResponse{}
			json.Unmarshal([]byte(rec.Body.String()), &res)

			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
			assert.Equal(t, "Incorrect email or password!", res.Message)
		}
	})
}
