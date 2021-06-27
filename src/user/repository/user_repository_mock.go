package repository

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (r *UserRepositoryMock) FindById(id string) (*user.User, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), nil
}

func (r *UserRepositoryMock) FindOne(query interface{}, data ...interface{}) (*user.User, error) {
	args := r.Called(query, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), nil
}
