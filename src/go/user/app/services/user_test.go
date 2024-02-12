package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	mocksmodels "market.goravel.dev/user/app/mocks/models"
	"market.goravel.dev/user/app/models"
)

type UserTestSuite struct {
	suite.Suite
	mockUser *mocksmodels.UserInterface
	userImpl *UserImpl
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupTest() {
	s.mockUser = new(mocksmodels.UserInterface)
	s.userImpl = &UserImpl{
		userModel: s.mockUser,
	}
}

func (s *UserTestSuite) TestIsEmailExist() {
	var (
		email = "hello@goravel.dev"
	)

	tests := []struct {
		name          string
		setup         func()
		expectedExist bool
		expectedErr   error
	}{
		{
			name: "Happy path",
			setup: func() {
				s.mockUser.On("GetUserByEmail", email, []string{"id"}).Return(&models.User{
					UUIDModel: models.UUIDModel{ID: 1},
				}, nil).Once()
			},
			expectedExist: true,
		},
		{
			name: "Sad path - GetUserByEmail returns error",
			setup: func() {
				s.mockUser.On("GetUserByEmail", email, []string{"id"}).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			exist, err := s.userImpl.IsEmailExist(email)
			s.Equal(test.expectedExist, exist)
			s.Equal(test.expectedErr, err)

			s.mockUser.AssertExpectations(s.T())
		})
	}
}
