package usecase

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/url"
)

type urlUsecase struct {
	urlRepository url.IUrlRepository
}

func NewUrlUsecase(urlRepository url.IUrlRepository) url.IUrlUsecase {
	return &urlUsecase{
		urlRepository,
	}
}

func (u *urlUsecase) FindById(id string) (*url.Url, error) {
	result, err := u.urlRepository.FindById(id)
	return result, err
}

func (u *urlUsecase) Create(payload *url.CreateUrlDto) (*url.Url, error) {
	result, err := u.urlRepository.Create(payload)
	return result, err
}
