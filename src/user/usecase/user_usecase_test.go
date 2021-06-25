package usecase

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user/repository"
	"github.com/stretchr/testify/assert"
)

func Test_User_Usecase_FindById(t *testing.T) {
	mockUserRepo := &repository.UserRepositoryMock{}
	now, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Run("Should return user data without password if found", func(t *testing.T) {
		id, _ := uuid.NewV4()

		returnValue := &user.User{
			ID:        id,
			Email:     "test@ranggarifqi.com",
			Password:  "asdsadsadasdsad",
			FullName:  "Test",
			CreatedAt: now,
			UpdatedAt: now,
		}
		mockUserRepo.On("FindById", id.String()).Return(returnValue, nil)

		usecase := NewUserUsecase(mockUserRepo)

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
}
