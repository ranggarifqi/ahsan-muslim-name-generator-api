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
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), nil
}
