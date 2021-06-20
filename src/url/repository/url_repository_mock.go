package repository

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/url"
	"github.com/stretchr/testify/mock"
)

type UrlRepositoryMock struct {
	mock.Mock
}

func (r *UrlRepositoryMock) FindById(id string) (*url.Url, error) {
	args := r.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*url.Url), nil
}

func (r *UrlRepositoryMock) Create(payload *url.CreateUrlDto) (*url.Url, error) {
	args := r.Called(payload)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*url.Url), nil
}
