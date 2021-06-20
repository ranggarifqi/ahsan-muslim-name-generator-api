package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/ranggarifqi/qashir-api/src/url"
	"github.com/ranggarifqi/qashir-api/src/url/repository"
	"github.com/stretchr/testify/assert"
)

func Test_Usecase_Url_FindById(t *testing.T) {
	t.Run("Should return something if found", func(t *testing.T) {
		repository := &repository.UrlRepositoryMock{}

		now, _ := time.Parse("2006-01-02", "2021-01-01")

		expected := &url.Url{
			ID:        "testid",
			Url:       "https://ranggarifqi.com",
			CreatedAt: now,
		}

		repository.On("FindById", "testid").Return(expected, nil)

		usecase := NewUrlUsecase(repository)

		result, err := usecase.FindById("testid")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("Should return error if not found", func(t *testing.T) {
		repository := &repository.UrlRepositoryMock{}

		err := errors.New("Error: Not found")
		repository.On("FindById", "testid").Return(nil, err)

		usecase := NewUrlUsecase(repository)

		result, err := usecase.FindById("testid")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "Error: Not found", err.Error())
	})
}

func Test_Usecase_Url_Create(t *testing.T) {
	t.Run("Should return Url object if successful", func(t *testing.T) {
		repository := &repository.UrlRepositoryMock{}

		now, _ := time.Parse("2006-01-02", "2021-01-01")
		expected := &url.Url{
			ID:        "testid",
			Url:       "https://ranggarifqi.com",
			CreatedAt: now,
		}

		input := &url.CreateUrlDto{
			Url: "https://ranggarifqi.com",
		}

		repository.On("Create", input).Return(expected, nil)

		usecase := NewUrlUsecase(repository)

		result, err := usecase.Create(input)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})
}
