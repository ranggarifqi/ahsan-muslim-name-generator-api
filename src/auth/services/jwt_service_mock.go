package service

import (
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
	"github.com/stretchr/testify/mock"
)

type JwtServiceMock struct {
	mock.Mock
}

func (s *JwtServiceMock) GetAuthToken(claim *auth.AuthClaim) (*string, error) {
	args := s.Called(claim)
	return args.Get(0).(*string), args.Error(1)
}
