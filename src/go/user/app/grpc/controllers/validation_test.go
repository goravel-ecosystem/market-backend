package controllers

import (
	"context"
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
)

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
			name:      "Happy path",
			email:     "hello@goravel.dev",
			setup:     func() {},
			expectErr: nil,
		},
		{
			name:  "Sad path - email invalid",
			email: "hello@goravel",
			setup: func() {
				mockLang.On("Get", "invalid.email").Return("invalid email").Once()
			},
			expectErr: errors.New("invalid email"),
		},
		{
			name:  "Sad path - email is empty",
			email: "",
			setup: func() {
				mockLang.On("Get", "required.email").Return("email is required").Once()
			},
			expectErr: errors.New("email is required"),
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
