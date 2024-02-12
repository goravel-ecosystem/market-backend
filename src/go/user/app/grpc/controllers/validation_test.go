package controllers

import (
	"context"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"

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
			expectErr: utilserrors.NewValidate("required email"),
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
			expectErr: utilserrors.NewValidate("password is required"),
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
			expectErr: utilserrors.NewValidate("required email"),
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
			expectErr: utilserrors.NewValidate("required name"),
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
			expectErr: utilserrors.NewValidate("required password"),
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
			expectErr: utilserrors.NewValidate("invalid password min"),
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
			expectErr: utilserrors.NewValidate("required code"),
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
			expectErr: utilserrors.NewValidate("required code_key"),
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
			expectErr: utilserrors.NewValidate("invalid email"),
		},
		{
			name:  "Sad path - email is empty",
			email: "",
			setup: func() {
				mockLang.On("Get", "required.email").Return("required email").Once()
			},
			expectErr: utilserrors.NewValidate("required email"),
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
