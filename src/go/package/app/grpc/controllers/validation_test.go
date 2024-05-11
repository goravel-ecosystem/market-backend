package controllers

import (
	"context"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	"github.com/goravel/framework/support/str"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	protopackage "market.goravel.dev/proto/package"
	utilserrors "market.goravel.dev/utils/errors"
)

func TestValidateCreatePackageRequest(t *testing.T) {
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
		request   *protopackage.CreatePackageRequest
		setup     func()
		expectErr error
	}{
		{
			name: "Happy path",
			request: &protopackage.CreatePackageRequest{
				UserId:        "1",
				Name:          "krishan",
				Cover:         "https://goravel.dev/covert.jpg",
				Summary:       "summary",
				Description:   "description",
				Url:           "https://goravel.dev",
				Tags:          []string{"tag1", "tag2"},
				IsPublic:      1,
				Version:       "1.0.0",
				LastUpdatedAt: "2021-09-01T00:00:00Z",
			},
			setup: func() {},
		},
		{
			name: "Empty user id",
			request: &protopackage.CreatePackageRequest{
				UserId: "",
				Name:   "krishan",
				Url:    "https://goravel.dev",
			},
			setup: func() {
				mockLang.On("Get", "required.user_id").Return("user id is required")
			},
			expectErr: utilserrors.NewBadRequest("user id is required"),
		},
		{
			name: "Empty name",
			request: &protopackage.CreatePackageRequest{
				UserId: "1",
				Name:   "",
				Url:    "https://goravel.dev",
			},
			setup: func() {
				mockLang.On("Get", "required.name").Return("name is required").Once()
			},
			expectErr: utilserrors.NewBadRequest("name is required"),
		},
		{
			name: "Name is too long",
			request: &protopackage.CreatePackageRequest{
				UserId: "1",
				Name:   str.Of("Krishan").Repeat(20).String(),
				Url:    "https://goravel.dev",
			},
			setup: func() {
				mockLang.On("Get", "max.name", mock.Anything).Return("Name must be less than 100").Once()
			},
			expectErr: utilserrors.NewBadRequest("Name must be less than 100"),
		},
		{
			name: "Empty url",
			request: &protopackage.CreatePackageRequest{
				UserId: "1",
				Name:   "krishan",
				Url:    "",
			},
			setup: func() {
				mockLang.On("Get", "required.url").Return("url is required").Once()
			},
			expectErr: utilserrors.NewBadRequest("url is required"),
		},
		{
			name: "Url is too long",
			request: &protopackage.CreatePackageRequest{
				UserId: "1",
				Name:   "krishan",
				Url:    str.Of("https://goravel.dev").Append(str.Of("/").Repeat(100).String()).String(),
			},
			setup: func() {
				mockLang.On("Get", "max.url", mock.Anything).Return("Url must be less than 100").Once()
			},
			expectErr: utilserrors.NewBadRequest("Url must be less than 100"),
		},
		{
			name: "Too many tags",
			request: &protopackage.CreatePackageRequest{
				UserId: "1",
				Name:   "krishan",
				Url:    "https://goravel.dev",
				Tags:   []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10", "tag11"},
			},
			setup: func() {
				mockLang.On("Get", "max.tags", mock.Anything).Return("Tags must be less than 10").Once()
			},
			expectErr: utilserrors.NewBadRequest("Tags must be less than 10"),
		},
		{
			name: "Summary is too long",
			request: &protopackage.CreatePackageRequest{
				UserId:  "1",
				Name:    "krishan",
				Url:     "https://goravel.dev",
				Summary: str.Of("summary").Repeat(30).String(),
			},
			setup: func() {
				mockLang.On("Get", "max.summary", mock.Anything).Return("Summary must be less than 200").Once()
			},
			expectErr: utilserrors.NewBadRequest("Summary must be less than 200"),
		},
		{
			name: "Description is too long",
			request: &protopackage.CreatePackageRequest{
				UserId:      "1",
				Name:        "krishan",
				Url:         "https://goravel.dev",
				Description: str.Of("description").Repeat(1000).String(),
			},
			setup: func() {
				mockLang.On("Get", "max.description", mock.Anything).Return("Description must be less than 10000").Once()
			},
			expectErr: utilserrors.NewBadRequest("Description must be less than 10000"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()
			assert.Equal(t, test.expectErr, validateCreatePackageRequest(ctx, test.request))

			mockLang.AssertExpectations(t)
		})
	}
}
