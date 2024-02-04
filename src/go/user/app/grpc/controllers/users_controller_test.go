package controllers

import (
	"context"
	"errors"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"

	mocksservice "market.goravel.dev/user/app/mocks/services"
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
	mockFactory := testingmock.Factory()
	s.mockLang = mockFactory.Lang(s.ctx)
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
			name: "Sad path - SendEmailRegisterCode returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: email,
			},
			setup: func() {
				s.mockNotificationService.On("SendEmailRegisterCode", s.ctx, email).Return("", errors.New("error")).Once()
			},
			expectedErr: utilserrors.NewValidate("error"),
		},
		{
			name: "Sad path - validateGetEmailRegisterCodeRequest returns error",
			request: &protouser.GetEmailRegisterCodeRequest{
				Email: "",
			},
			setup: func() {
				s.mockLang.On("Get", "required.email").Return("email is required").Once()
			},
			expectedErr: utilserrors.NewValidate("email is required"),
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
