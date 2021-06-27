package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper/testutil"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/response"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	userRepo "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	userUC "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/usecase"
	"github.com/stretchr/testify/assert"
)

type Mocks struct {
	mockUserRepository *userRepo.UserRepositoryMock
}

func setupMock() Mocks {
	return Mocks{
		mockUserRepository: &userRepo.UserRepositoryMock{},
	}
}

func Test_User_handler_FindById(t *testing.T) {
	t.Run("Should return correct response if ID is found", func(t *testing.T) {
		id, _ := uuid.NewV4()

		foundUser := &user.User{
			ID:       id,
			Email:    "test@ranggarifqi.com",
			Password: "encryptedpassword",
			FullName: "Test",
		}

		/* Setup Mocks */
		mocks := setupMock()
		mocks.mockUserRepository.On("FindById", id.String()).Return(foundUser, nil)

		/* Setup Handler */
		userUsecase := userUC.NewUserUsecase(mocks.mockUserRepository)
		handler := UserHandler{
			uc: userUsecase,
		}

		/* Setup request */
		e := testutil.SetupServer()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/users/:id", nil)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues(id.String())

		/* Assertions */
		if assert.NoError(t, handler.FindById(ctx)) {
			res := response.SuccessResponse{}
			json.Unmarshal([]byte(rec.Body.String()), &res)
			resData, _ := res.Data.(map[string]interface{})

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Equal(t, "Data fetched successfully", res.Message)
			assert.Equal(t, id.String(), resData["id"])
		}
	})
}
