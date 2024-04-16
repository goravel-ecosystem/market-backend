package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/goravel/framework/http"
	mocksauth "github.com/goravel/framework/mocks/auth"
	mockshash "github.com/goravel/framework/mocks/hash"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	protouser "market.goravel.dev/proto/user"
	mocksservice "market.goravel.dev/user/app/mocks/services"
	"market.goravel.dev/user/app/models"
	utilserrors "market.goravel.dev/utils/errors"
	utilsresponse "market.goravel.dev/utils/response"
)

type UsersControllerSuite struct {
	suite.Suite
	ctx                     context.Context
	userController          *UserController
	mockAuth                *mocksauth.Auth
	mockHash                *mockshash.Hash
	mockLang                *mockstranslation.Translator
	mockNotificationService *mocksservice.Notification
	mockUserService         *mocksservice.User
}

func TestUsersControllerSuite(t *testing.T) {
	suite.Run(t, new(UsersControllerSuite))
}

func (s *UsersControllerSuite) SetupTest() {
	s.ctx = context.Background()
	mockFactory := testingmock.Factory()
	s.mockAuth = mockFactory.Auth(http.Background())
	s.mockHash = mockFactory.Hash()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.mockNotificationService = &mocksservice.Notification{}
	s.mockUserService = &mocksservice.User{}
	s.userController = &UserController{
		notificationService: s.mockNotificationService,
		userService:         s.mockUserService,
	}
}

func (s *UsersControllerSuite) TestEmailLogin() {
	var (
		email          = "hello@goravel.dev"
		password       = "password"
		hashedPassword = "hashed_password"

		user = models.User{
			UUIDModel: models.UUIDModel{
				ID: 1,
			},
			Email:    email,
			Password: hashedPassword,
		}
	)

	tests := []struct {
		name             string
		request          *protouser.EmailLoginRequest
		setup            func()
		expectedResponse *protouser.EmailLoginResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protouser.EmailLoginRequest{
				Email:    email,
				Password: password,
			},
			setup: func() {
				s.mockUserService.On("GetUserByEmail", email).Return(&user, nil).Once()
				s.mockHash.On("Check", password, hashedPassword).Return(true).Once()
				s.mockAuth.On("LoginUsingID", user.ID).Return("token", nil).Once()
			},
			expectedResponse: &protouser.EmailLoginResponse{
				Status: utilsresponse.NewOkStatus(),
				User:   user.ToProto(),
				Token:  "Bearer token",
			},
		},
		{
			name: "Sad path - email is invalid",
			request: &protouser.EmailLoginRequest{
				Email:    "",
				Password: password,
			},
			setup: func() {
				s.mockLang.On("Get", "required.email").Return("required email").Once()
			},
			expectedErr: utilserrors.NewBadRequest("required email"),
		},
		{
			name: "Sad path - password is invalid",
			request: &protouser.EmailLoginRequest{
				Email:    email,
				Password: "",
			},
			setup: func() {
				s.mockLang.On("Get", "required.password").Return("required password").Once()
			},
			expectedErr: utilserrors.NewBadRequest("required password"),
		},
		{
			name: "Sad path - GetUserByEmail returns error",
			request: &protouser.EmailLoginRequest{
				Email:    email,
				Password: password,
			},
			setup: func() {
				s.mockUserService.On("GetUserByEmail", email).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - password is wrong",
			request: &protouser.EmailLoginRequest{
				Email:    email,
				Password: password,
			},
			setup: func() {
				s.mockUserService.On("GetUserByEmail", email).Return(&user, nil).Once()
				s.mockHash.On("Check", password, hashedPassword).Return(false).Once()
				s.mockLang.On("Get", "invalid.password.error").Return("invalid password error").Once()
			},
			expectedErr: utilserrors.NewBadRequest("invalid password error"),
		},
		{
			name: "Sad path - LoginUsingID returns error",
			request: &protouser.EmailLoginRequest{
				Email:    email,
				Password: password,
			},
			setup: func() {
				s.mockUserService.On("GetUserByEmail", email).Return(&user, nil).Once()
				s.mockHash.On("Check", password, hashedPassword).Return(true).Once()
				s.mockAuth.On("LoginUsingID", user.ID).Return("", errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.userController.EmailLogin(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockAuth.AssertExpectations(s.T())
			s.mockHash.AssertExpectations(s.T())
			s.mockLang.AssertExpectations(s.T())
			s.mockUserService.AssertExpectations(s.T())
		})
	}
}

func (s *UsersControllerSuite) TestEmailRegister() {
	var (
		code     = "code"
		codeKey  = "code_key"
		email    = "hello@goravel.dev"
		name     = "name"
		password = "password"

		user = models.User{
			UUIDModel: models.UUIDModel{
				ID: 1,
			},
			Name:  name,
			Email: email,
		}
	)

	tests := []struct {
		name             string
		request          *protouser.EmailRegisterRequest
		setup            func()
		expectedResponse *protouser.EmailRegisterResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("VerifyEmailRegisterCode", codeKey, code).Return(true).Once()
				s.mockUserService.On("Register", name, email, password).Return(&user, nil).Once()
				s.mockAuth.On("LoginUsingID", user.ID).Return("token", nil).Once()
			},
			expectedResponse: &protouser.EmailRegisterResponse{
				Status: utilsresponse.NewOkStatus(),
				User:   user.ToProto(),
				Token:  "Bearer token",
			},
		},
		{
			name: "Sad path - LoginUsingID returns error",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("VerifyEmailRegisterCode", codeKey, code).Return(true).Once()
				s.mockUserService.On("Register", name, email, password).Return(&user, nil).Once()
				s.mockAuth.On("LoginUsingID", user.ID).Return("", errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - Register returns error",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("VerifyEmailRegisterCode", codeKey, code).Return(true).Once()
				s.mockUserService.On("Register", name, email, password).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - VerifyEmailRegisterCode returns false",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("VerifyEmailRegisterCode", codeKey, code).Return(false).Once()
				s.mockLang.On("Get", "invalid.code").Return("invalid code").Once()
			},
			expectedErr: utilserrors.NewBadRequest("invalid code"),
		},
		{
			name: "Sad path - IsEmailExist returns error",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - email already exist",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(true, nil).Once()
				s.mockLang.On("Get", "exist.email").Return("exist email").Once()
			},
			expectedErr: utilserrors.NewBadRequest("exist email"),
		},
		{
			name: "Sad path - email is empty",
			request: &protouser.EmailRegisterRequest{
				Name:     name,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "required.email").Return("email is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("email is required"),
		},
		{
			name: "Sad path - email is invalid",
			request: &protouser.EmailRegisterRequest{
				Email:    "email",
				Name:     name,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "invalid.email").Return("email is invalid").Once()
			},
			expectedErr: utilserrors.NewBadRequest("email is invalid"),
		},
		{
			name: "Sad path - name is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    email,
				Password: password,
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "required.name").Return("name is empty").Once()
			},
			expectedErr: utilserrors.NewBadRequest("name is empty"),
		},
		{
			name: "Sad path - password is empty",
			request: &protouser.EmailRegisterRequest{
				Email:   email,
				Name:    name,
				Code:    code,
				CodeKey: codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "required.password").Return("password is empty").Once()
			},
			expectedErr: utilserrors.NewBadRequest("password is empty"),
		},
		{
			name: "Sad path - password len < 6",
			request: &protouser.EmailRegisterRequest{
				Email:    email,
				Name:     name,
				Password: "123",
				Code:     code,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "invalid.password.min").Return("password is invalid").Once()
			},
			expectedErr: utilserrors.NewBadRequest("password is invalid"),
		},
		{
			name: "Sad path - code is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    email,
				Name:     name,
				Password: password,
				CodeKey:  codeKey,
			},
			setup: func() {
				s.mockLang.On("Get", "required.code").Return("code is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("code is required"),
		},
		{
			name: "Sad path - code key is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    email,
				Name:     name,
				Password: password,
				Code:     code,
			},
			setup: func() {
				s.mockLang.On("Get", "required.code_key").Return("code key is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("code key is required"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.userController.EmailRegister(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockAuth.AssertExpectations(s.T())
			s.mockLang.AssertExpectations(s.T())
			s.mockNotificationService.AssertExpectations(s.T())
			s.mockUserService.AssertExpectations(s.T())
		})
	}
}

func (s *UsersControllerSuite) TestGetEmailRegisterCode() {
	var (
		email = "hello@goravel.dev"
		key   = "key"
	)

	tests := []struct {
		name             string
		request          *protouser.GetEmailRegisterCodeRequest
		setup            func()
		expectedResponse *protouser.GetEmailRegisterCodeResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("SendEmailRegisterCode", s.ctx, email).Return(key, nil).Once()
			},
			expectedResponse: &protouser.GetEmailRegisterCodeResponse{
				Status: utilsresponse.NewOkStatus(),
				Key:    key,
			},
		},
		{
			name: "Sad path - SendEmailRegisterCode returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, nil).Once()
				s.mockNotificationService.On("SendEmailRegisterCode", s.ctx, email).Return("", errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - IsEmailExist returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(false, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - mail already exist",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockUserService.On("IsEmailExist", email).Return(true, nil).Once()
				s.mockLang.On("Get", "exist.email").Return("email already exist").Once()
			},
			expectedErr: utilserrors.NewBadRequest("email already exist"),
		},
		{
			name: "Sad path - validateGetEmailRegisterCodeRequest returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: "",
			},
			setup: func() {
				s.mockLang.On("Get", "required.email").Return("email is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("email is required"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.userController.GetEmailRegisterCode(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockNotificationService.AssertExpectations(s.T())
			s.mockUserService.AssertExpectations(s.T())
		})
	}
}

func (s *UsersControllerSuite) TestGetUser() {
	var (
		userID = "1"
		name   = "Goravel"
	)

	tests := []struct {
		name             string
		request          *protouser.GetUserRequest
		setup            func()
		expectedResponse *protouser.GetUserResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protouser.GetUserRequest{
				UserId: userID,
			},
			setup: func() {
				s.mockUserService.On("GetUserByID", userID).Return(&models.User{
					UUIDModel: models.UUIDModel{
						ID: 1,
					},
					Name: name,
				}, nil).Once()
			},
			expectedResponse: &protouser.GetUserResponse{
				Status: utilsresponse.NewOkStatus(),
				User: &protouser.User{
					Id:   "1",
					Name: name,
				},
			},
		},
		{
			name: "Sad path - user id is empty",
			request: &protouser.GetUserRequest{
				UserId: "",
			},
			setup: func() {
				s.mockLang.On("Get", "required.user_id").Return("required user id").Once()
			},
			expectedErr: utilserrors.NewBadRequest("required user id"),
		},
		{
			name: "Sad path - GetUserByID returns error",
			request: &protouser.GetUserRequest{
				UserId: userID,
			},
			setup: func() {
				s.mockUserService.On("GetUserByID", userID).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - user not found",
			request: &protouser.GetUserRequest{
				UserId: userID,
			},
			setup: func() {
				s.mockUserService.On("GetUserByID", userID).Return(&models.User{}, nil).Once()
				s.mockLang.On("Get", "not_exist.user").Return("User not found").Once()
			},
			expectedErr: utilserrors.NewNotFound("User not found"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.userController.GetUser(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockAuth.AssertExpectations(s.T())
			s.mockHash.AssertExpectations(s.T())
			s.mockLang.AssertExpectations(s.T())
			s.mockUserService.AssertExpectations(s.T())
		})
	}

}
