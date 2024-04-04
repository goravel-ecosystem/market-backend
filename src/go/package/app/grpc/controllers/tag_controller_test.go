package controllers

import (
	"context"
	"errors"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	"market.goravel.dev/package/app/models"
	protopackage "market.goravel.dev/proto/package"
	utilsresponse "market.goravel.dev/utils/response"

	mocksservice "market.goravel.dev/package/app/mocks/services"
)

type TagControllerSuite struct {
	suite.Suite
	ctx            context.Context
	tagController  *TagController
	mockLang       *mockstranslation.Translator
	mockTagService *mocksservice.Tag
}

func TestTagControllerSuite(t *testing.T) {
	suite.Run(t, new(TagControllerSuite))
}

func (s *TagControllerSuite) SetupTest() {
	s.ctx = context.Background()
	mockFactory := testingmock.Factory()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.mockTagService = &mocksservice.Tag{}
	s.tagController = &TagController{
		tagService: s.mockTagService,
	}
}

func (s *TagControllerSuite) TestGetTags() {
	var (
		packageID = "1"
		userID    = "1"
		name      = "go"
	)

	tests := []struct {
		name             string
		request          *protopackage.GetTagRequest
		setup            func()
		expectedResponse *protopackage.GetTagResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.GetTagRequest{
				PackageId: packageID,
				Name:      name,
			},
			setup: func() {
				s.mockTagService.On("GetTags", packageID, "", name).Return([]*models.Tag{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: 1,
						Name:   "goravel",
					},
				}, nil).Once()
			},
			expectedResponse: &protopackage.GetTagResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags: []*protopackage.Tag{
					{
						Id:     "1",
						UserId: userID,
						Name:   "goravel",
					},
				},
			},
		},
		{
			name: "Sad path - GetTags returns error",
			request: &protopackage.GetTagRequest{
				Name: name,
			},
			setup: func() {
				s.mockTagService.On("GetTags", "", "", name).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - tags is empty",
			request: &protopackage.GetTagRequest{
				Name: name,
			},
			setup: func() {
				s.mockTagService.On("GetTags", "", "", name).Return([]*models.Tag{}, nil).Once()
				s.mockLang.On("Get", "not_exist.tags").Return("tags not found").Once()
			},
			expectedResponse: &protopackage.GetTagResponse{
				Status: utilsresponse.NewNotFoundStatus(errors.New("tags not found")),
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.tagController.GetTags(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockTagService.AssertExpectations(s.T())
		})
	}
}
