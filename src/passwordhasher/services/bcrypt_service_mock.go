package service

import "github.com/stretchr/testify/mock"

type BCryptServiceMock struct {
	mock.Mock
}

func (s *BCryptServiceMock) HashPassword(password string) (string, error) {
	args := s.Called(password)
	return args.String(0), args.Error(1)
}

func (s *BCryptServiceMock) ComparePassword(hashedPassword string, password string) (bool, error) {
	args := s.Called(hashedPassword, password)
	return args.Bool(0), args.Error(1)
}
