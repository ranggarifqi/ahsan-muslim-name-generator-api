package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	"github.com/stretchr/testify/assert"
)

type Mocks struct {
	mockUserRepository *repository.UserRepositoryMock
}

func setupMocks() Mocks {
	return Mocks{
		mockUserRepository: &repository.UserRepositoryMock{},
	}
}

func Test_User_Usecase_FindById(t *testing.T) {
	now, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Run("Should return user data without password if found", func(t *testing.T) {
		id, _ := uuid.NewV4()

		/* Setup Mocks */
		mocks := setupMocks()

		returnValue := &user.User{
			ID:        id,
			Email:     "test@ranggarifqi.com",
			Password:  "asdsadsadasdsad",
			FullName:  "Test",
			CreatedAt: now,
			UpdatedAt: now,
		}
		mocks.mockUserRepository.On("FindById", id.String()).Return(returnValue, nil)

		usecase := NewUserUsecase(mocks.mockUserRepository)

		result, err := usecase.FindById(id.String())

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &user.UserWithoutPassword{
			ID:        id,
			Email:     "test@ranggarifqi.com",
			FullName:  "Test",
			CreatedAt: now,
			UpdatedAt: now,
		}, result)
	})

	t.Run("Should return record not found error if ID not found", func(t *testing.T) {
		randomID, _ := uuid.NewV4()

		/* Setup mocks */
		mocks := setupMocks()
		mocks.mockUserRepository.On("FindById", randomID.String()).Return(nil, errors.New("record not found"))

		/* Setup usecase */
		usecase := NewUserUsecase(mocks.mockUserRepository)

		/* Assertion */
		result, err := usecase.FindById(randomID.String())

		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "record not found")
	})
}
