package controllers

import (
	"context"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	"github.com/goravel/framework/support/str"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

func TestValidateEmailLoginRequest(t *testing.T) {
	var (
		ctx      = context.Background()
		mockLang *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockLang = mockFactory.Lang(ctx)
	}

	tests := []struct {
		name      string
		request   *protouser.EmailLoginRequest
		setup     func()
		expectErr error
	}{
		{
			name: "Happy path",
			request: &protouser.EmailLoginRequest{
				Email:    "hello@goravel.com",
				Password: "password",
			},
			setup: func() {},
		},
		{
			name: "Sad path - email invalid",
			request: &protouser.EmailLoginRequest{
				Email:    "",
				Password: "password",
			},
			setup: func() {
				mockLang.On("Get", "required.email").Return("required email").Once()
			},
			expectErr: utilserrors.NewBadRequest("required email"),
		},
		{
			name: "Sad path - password is empty",
			request: &protouser.EmailLoginRequest{
				Email:    "hello@goravel.com",
				Password: "",
			},
			setup: func() {
				mockLang.On("Get", "required.password").Return("password is required").Once()
			},
			expectErr: utilserrors.NewBadRequest("password is required"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()
			assert.Equal(t, test.expectErr, validateEmailLoginRequest(ctx, test.request))

			mockLang.AssertExpectations(t)
		})
	}
}

func TestValidateEmailRegisterRequest(t *testing.T) {
	var (
		ctx      = context.Background()
		mockLang *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockLang = mockFactory.Lang(ctx)
	}

	tests := []struct {
		name      string
		request   *protouser.EmailRegisterRequest
		setup     func()
		expectErr error
	}{
		{
			name: "Happy path",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "Goravel",
				Password: "password",
				Code:     "123456",
				CodeKey:  "key",
			},
			setup: func() {},
		},
		{
			name: "Sad path - email invalid",
			request: &protouser.EmailRegisterRequest{
				Email:    "",
				Name:     "Goravel",
				Password: "password",
				Code:     "123456",
				CodeKey:  "key",
			},
			setup: func() {
				mockLang.On("Get", "required.email").Return("required email").Once()
			},
			expectErr: utilserrors.NewBadRequest("required email"),
		},
		{
			name: "Sad path - name is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "",
				Password: "password",
				Code:     "123456",
				CodeKey:  "key",
			},
			setup: func() {
				mockLang.On("Get", "required.name").Return("required name").Once()
			},
			expectErr: utilserrors.NewBadRequest("required name"),
		},
		{
			name: "Sad path - password is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "Goravel",
				Password: "",
				Code:     "123456",
				CodeKey:  "key",
			},
			setup: func() {
				mockLang.On("Get", "required.password").Return("required password").Once()
			},
			expectErr: utilserrors.NewBadRequest("required password"),
		},
		{
			name: "Sad path - the password is less than 6 characters",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "Goravel",
				Password: "123",
				Code:     "123456",
				CodeKey:  "key",
			},
			setup: func() {
				mockLang.On("Get", "invalid.password.min").Return("invalid password min").Once()
			},
			expectErr: utilserrors.NewBadRequest("invalid password min"),
		},
		{
			name: "Sad path - code is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "Goravel",
				Password: "password",
				Code:     "",
				CodeKey:  "key",
			},
			setup: func() {
				mockLang.On("Get", "required.code").Return("required code").Once()
			},
			expectErr: utilserrors.NewBadRequest("required code"),
		},
		{
			name: "Sad path - code key is empty",
			request: &protouser.EmailRegisterRequest{
				Email:    "hello@goravel.com",
				Name:     "Goravel",
				Password: "password",
				Code:     "123456",
				CodeKey:  "",
			},
			setup: func() {
				mockLang.On("Get", "required.code_key").Return("required code_key").Once()
			},
			expectErr: utilserrors.NewBadRequest("required code_key"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()
			assert.Equal(t, test.expectErr, validateEmailRegisterRequest(ctx, test.request))

			mockLang.AssertExpectations(t)
		})
	}
}

func TestValidateEmailValid(t *testing.T) {
	var (
		ctx      = context.Background()
		mockLang *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockLang = mockFactory.Lang(ctx)
	}

	tests := []struct {
		name      string
		email     string
		setup     func()
		expectErr error
	}{
		{
			name:  "Happy path",
			email: "hello@goravel.dev",
			setup: func() {},
		},
		{
			name:  "Sad path - email invalid",
			email: "hello@goravel",
			setup: func() {
				mockLang.On("Get", "invalid.email").Return("invalid email").Once()
			},
			expectErr: utilserrors.NewBadRequest("invalid email"),
		},
		{
			name:  "Sad path - email is empty",
			email: "",
			setup: func() {
				mockLang.On("Get", "required.email").Return("required email").Once()
			},
			expectErr: utilserrors.NewBadRequest("required email"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()
			assert.Equal(t, test.expectErr, validateEmailValid(ctx, test.email))

			mockLang.AssertExpectations(t)
		})
	}
}

func TestValidateUpdateUserRequest(t *testing.T) {
	var (
		ctx      = context.Background()
		id       = "1"
		userID   = "1"
		name     = "krishan"
		summary  = "I am a developer"
		password = "password"

		mockLang *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockLang = mockFactory.Lang(ctx)
	}

	tests := []struct {
		name      string
		request   *protouser.UpdateUserRequest
		setup     func()
		expectErr error
	}{
		{
			name: "Happy path",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Summary:  summary,
				Password: password,
			},
			setup: func() {},
		},
		{
			name: "Empty id",
			request: &protouser.UpdateUserRequest{
				Id:     "",
				UserId: userID,
				Name:   name,
			},
			setup: func() {
				mockLang.On("Get", "required.id").Return("id is required")
			},
			expectErr: utilserrors.NewBadRequest("id is required"),
		},
		{
			name: "Empty user id",
			request: &protouser.UpdateUserRequest{
				Id:     id,
				UserId: "",
				Name:   name,
			},
			setup: func() {
				mockLang.On("Get", "required.user_id").Return("user id is required")
			},
			expectErr: utilserrors.NewBadRequest("user id is required"),
		},
		{
			name: "Empty name",
			request: &protouser.UpdateUserRequest{
				Id:     id,
				UserId: userID,
				Name:   "",
			},
			setup: func() {
				mockLang.On("Get", "required.name").Return("name is required").Once()
			},
			expectErr: utilserrors.NewBadRequest("name is required"),
		},
		{
			name: "Name is too long",
			request: &protouser.UpdateUserRequest{
				Id:     id,
				UserId: userID,
				Name:   str.Of("Krishan").Repeat(20).String(),
			},
			setup: func() {
				mockLang.On("Get", "invalid.name.max", mock.Anything).Return("name must be less than 50").Once()
			},
			expectErr: utilserrors.NewBadRequest("name must be less than 50"),
		},
		{
			name: "Summary is too long",
			request: &protouser.UpdateUserRequest{
				Id:      id,
				UserId:  userID,
				Name:    name,
				Summary: str.Of(summary).Repeat(20).String(),
			},
			setup: func() {
				mockLang.On("Get", "invalid.summery.max", mock.Anything).Return("summary must be less than 200").Once()
			},
			expectErr: utilserrors.NewBadRequest("summary must be less than 200"),
		},
		{
			name: "Password is too short",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Password: "123",
			},
			setup: func() {
				mockLang.On("Get", "invalid.password.min").Return("password must be more than 6 characters").Once()
			},
			expectErr: utilserrors.NewBadRequest("password must be more than 6 characters"),
		},
		{
			name: "Password is too long",
			request: &protouser.UpdateUserRequest{
				Id:       id,
				UserId:   userID,
				Name:     name,
				Password: str.Of("password").Repeat(20).String(),
			},
			setup: func() {
				mockLang.On("Get", "invalid.password.max", mock.Anything).Return("password must be less than 50").Once()
			},
			expectErr: utilserrors.NewBadRequest("password must be less than 50"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()
			assert.Equal(t, test.expectErr, validateUpdateUserRequest(ctx, test.request))

			mockLang.AssertExpectations(t)
		})
	}
}
