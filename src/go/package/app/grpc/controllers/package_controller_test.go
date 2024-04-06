package controllers

import (
	"context"
	"errors"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	utilsresponse "market.goravel.dev/utils/response"

	mocksservice "market.goravel.dev/package/app/mocks/services"
)

type PackageControllerSuite struct {
	suite.Suite
	ctx               context.Context
	packageController *PackageController
	mockLang          *mockstranslation.Translator
	mockTagService    *mocksservice.Tag
}

func TestPackageControllerSuite(t *testing.T) {
	suite.Run(t, new(PackageControllerSuite))
}

func (s *PackageControllerSuite) SetupTest() {
	s.ctx = context.Background()
	mockFactory := testingmock.Factory()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.mockTagService = &mocksservice.Tag{}
	s.packageController = &PackageController{
		tagService: s.mockTagService,
	}
}

func (s *PackageControllerSuite) TestGetTags() {
	var (
		packageID  = "1"
		userID     = "1"
		name       = "go"
		pagination = &protobase.Pagination{
			Page:  1,
			Limit: 10,
		}
		total int64
	)

	tests := []struct {
		name             string
		request          *protopackage.GetTagsRequest
		setup            func()
		expectedResponse *protopackage.GetTagsResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.GetTagsRequest{
				Pagination: pagination,
				Query: &protopackage.TagsQuery{
					PackageId: packageID,
					Name:      name,
				},
			},
			setup: func() {
				s.mockTagService.On("GetTags", packageID, name, pagination, &total).Return([]*models.Tag{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: 1,
						Name:   "goravel",
					},
				}, nil).Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags: []*protopackage.Tag{
					{
						Id:     "1",
						UserId: userID,
						Name:   "goravel",
					},
				},
				Total: total,
			},
		},
		{
			name: "Sad path - GetTags returns error",
			request: &protopackage.GetTagsRequest{
				Pagination: pagination,
				Query: &protopackage.TagsQuery{
					Name: name,
				},
			},
			setup: func() {
				s.mockTagService.On("GetTags", "", name, pagination, &total).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - tags is empty",
			request: &protopackage.GetTagsRequest{
				Pagination: pagination,
				Query: &protopackage.TagsQuery{
					Name: name,
				},
			},
			setup: func() {
				s.mockTagService.On("GetTags", "", name, pagination, &total).Return([]*models.Tag{}, nil).Once()
				s.mockLang.On("Get", "not_exist.tags").Return("tags not found").Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewNotFoundStatus(errors.New("tags not found")),
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.packageController.GetTags(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockTagService.AssertExpectations(s.T())
		})
	}
}
