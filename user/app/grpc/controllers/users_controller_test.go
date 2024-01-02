package controllers

import (
	"context"
	"errors"
	"testing"

	protouser "github.com/goravel-ecosystem/market-backend/proto/user"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	mocksservice "github.com/goravel-ecosystem/market-backend/user/mocks/services"
)

type UsersControllerSuite struct {
	suite.Suite
	ctx                     context.Context
	usersController         *UsersController
	mockLang                *mockstranslation.Translator
	mockNotificationService *mocksservice.Notification
}

func TestUsersControllerSuite(t *testing.T) {
	suite.Run(t, new(UsersControllerSuite))
}

func (s *UsersControllerSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockLang = testingmock.Lang(s.ctx)
	s.mockNotificationService = &mocksservice.Notification{}
	s.usersController = &UsersController{
		notificationService: s.mockNotificationService,
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
				s.mockNotificationService.On("SendEmailRegisterCode", s.ctx, email).Return(key, nil).Once()
			},
			expectedResponse: &protouser.GetEmailRegisterCodeResponse{
				Status: NewOkStatus(),
				Key:    key,
			},
		},
		{
			name: "Error path - SendEmailRegisterCode returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockNotificationService.On("SendEmailRegisterCode", s.ctx, email).Return("", errors.New("error")).Once()
			},
			expectedResponse: &protouser.GetEmailRegisterCodeResponse{
				Status: NewBadRequestStatus(errors.New("error")),
			},
		},
		{
			name: "Error path - validateGetEmailRegisterCodeRequest returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: "",
			},
			setup: func() {
				s.mockLang.On("Get", "required.email").Return("required.email", nil).Once()
			},
			expectedResponse: &protouser.GetEmailRegisterCodeResponse{
				Status: NewBadRequestStatus(errors.New("required.email")),
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.usersController.GetEmailRegisterCode(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)
		})
	}
}
