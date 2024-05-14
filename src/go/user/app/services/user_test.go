package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	mockshash "github.com/goravel/framework/mocks/hash"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	protouser "market.goravel.dev/proto/user"
	mocksmodels "market.goravel.dev/user/app/mocks/models"
	"market.goravel.dev/user/app/models"
	utilserrors "market.goravel.dev/utils/errors"
)

type UserTestSuite struct {
	suite.Suite
	ctx      context.Context
	mockUser *mocksmodels.UserInterface
	mockLang *mockstranslation.Translator
	mockHash *mockshash.Hash
	userImpl *UserImpl
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupTest() {
	s.mockUser = new(mocksmodels.UserInterface)
	s.ctx = context.Background()
	mockFactory := testingmock.Factory()
	mockFactory.Log()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.mockHash = mockFactory.Hash()
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

func (s *UserTestSuite) TestUpdateUser() {
	var (
		id       = "1"
		userID   = "1"
		name     = "krishan"
		avatar   = "https://avatar.com/avatar.jpg"
		summary  = "I am a developer"
		password = "password"
	)

	tests := []struct {
		name        string
		request     *protouser.UpdateUserRequest
		setup       func()
		expectUser  *models.User
		expectedErr error
	}{
		{
			name: "Happy path - UpdateUser with ID",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Avatar:   avatar,
				Summary:  summary,
				Password: password,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(&models.User{UUIDModel: models.UUIDModel{ID: 1}}, nil).Once()
				s.mockHash.On("Make", password).Return(password, nil).Once()
				s.mockUser.On("UpdateUser", mock.MatchedBy(func(user *models.User) bool {
					return user.Name == name && user.Avatar == avatar && user.Summary == summary && user.Password != ""
				})).Return(nil).Once()
			},
			expectUser: &models.User{UUIDModel: models.UUIDModel{ID: 1}, Name: name, Password: password, Avatar: avatar, Summary: summary},
		},
		{
			name: "Happy path - GetUserByID returns error",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Avatar:   avatar,
				Summary:  summary,
				Password: password,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - Package does not exist",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Avatar:   avatar,
				Summary:  summary,
				Password: password,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(&models.User{UUIDModel: models.UUIDModel{ID: 0}}, nil).Once()
				s.mockLang.On("Get", "not_exist.user").Return("not_exist.user").Once()
			},
			expectedErr: utilserrors.NewNotFound("not_exist.user"),
		},
		{
			name: "Sad path - User ID does not match",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   "2",
				Name:     name,
				Avatar:   avatar,
				Summary:  summary,
				Password: password,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(&models.User{UUIDModel: models.UUIDModel{ID: 1}}, nil).Once()
				s.mockLang.On("Get", "forbidden.update_user").Return("forbidden.update_user").Once()
			},
			expectedErr: utilserrors.NewUnauthorized("forbidden.update_user"),
		},
		{
			name: "Sad path - UpdateUser returns error",
			request: &protouser.UpdateUserRequest{
				Id:      id,
				UserId:  userID,
				Name:    name,
				Avatar:  avatar,
				Summary: summary,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(&models.User{UUIDModel: models.UUIDModel{ID: 1}}, nil).Once()
				s.mockUser.On("UpdateUser", mock.MatchedBy(func(user *models.User) bool {
					return user.Name == name && user.Avatar == avatar && user.Summary == summary
				})).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - Password hashing error",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Avatar:   avatar,
				Password: password,
				Summary:  summary,
			},
			setup: func() {
				s.mockUser.On("GetUserByID", id, []string{}).Return(&models.User{UUIDModel: models.UUIDModel{ID: 1}}, nil).Once()
				s.mockHash.On("Make", password).Return("", errors.New("error")).Once()

			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()

			user, err := s.userImpl.UpdateUser(s.ctx, test.request)
			if test.expectedErr != nil {
				s.Nil(user)
				s.Equal(test.expectedErr, err)
			} else {
				s.Nil(err)
				s.Equal(test.expectUser, user)
			}

			s.mockUser.AssertExpectations(s.T())
			s.mockLang.AssertExpectations(s.T())
		})
	}
}
