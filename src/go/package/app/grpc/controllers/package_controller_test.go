package controllers

import (
	"context"
	"errors"
	"fmt"
	"testing"

	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/suite"

	mocksservice "market.goravel.dev/package/app/mocks/services"
	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	utilsresponse "market.goravel.dev/utils/response"
)

type PackageControllerSuite struct {
	suite.Suite
	ctx                context.Context
	packageController  *PackageController
	mockLang           *mockstranslation.Translator
	mockPackageService *mocksservice.Package
	mockTagService     *mocksservice.Tag
	mockUserService    *mocksservice.User
}

func TestPackageControllerSuite(t *testing.T) {
	suite.Run(t, new(PackageControllerSuite))
}

func (s *PackageControllerSuite) SetupTest() {
	s.ctx = context.Background()
	mockFactory := testingmock.Factory()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.mockPackageService = &mocksservice.Package{}
	s.mockTagService = &mocksservice.Tag{}
	s.mockUserService = &mocksservice.User{}
	s.packageController = &PackageController{
		packageService: s.mockPackageService,
		tagService:     s.mockTagService,
		userService:    s.mockUserService,
	}
}

func (s *PackageControllerSuite) TestGetPackage() {
	// TODO: implement me
}

func (s *PackageControllerSuite) TestGetTags() {
	var (
		packageID  = "1"
		userID     = uint64(1)
		name       = "goravel"
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
				total = 1
				s.mockTagService.On("GetTags", packageID, name, pagination).Return([]*models.Tag{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: userID,
						Name:   name,
					},
				}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags: []*protopackage.Tag{
					{
						Id:     "1",
						UserId: fmt.Sprint(userID),
						Name:   name,
					},
				},
				Total: 1,
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
				total = 0
				s.mockTagService.On("GetTags", "", name, pagination).Return(nil, total, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Happy path - tags is empty",
			request: &protopackage.GetTagsRequest{
				Pagination: pagination,
				Query: &protopackage.TagsQuery{
					Name: name,
				},
			},
			setup: func() {
				total = 0
				s.mockTagService.On("GetTags", "", name, pagination).Return([]*models.Tag{}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags:   []*protopackage.Tag{},
			},
		},
		{
			name: "Happy path - pagination is nil",
			request: &protopackage.GetTagsRequest{
				Query: &protopackage.TagsQuery{
					Name: name,
				},
			},
			setup: func() {
				total = 0
				s.mockTagService.On("GetTags", "", name, &protobase.Pagination{Page: 1, Limit: 10}).Return([]*models.Tag{}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags:   []*protopackage.Tag{},
			},
		},
		{
			name: "Happy path - query is nil",
			request: &protopackage.GetTagsRequest{
				Pagination: pagination,
			},
			setup: func() {
				total = 1
				s.mockTagService.On("GetTags", "", "", pagination).Return([]*models.Tag{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: userID,
						Name:   name,
					},
				}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetTagsResponse{
				Status: utilsresponse.NewOkStatus(),
				Tags: []*protopackage.Tag{
					{
						Id:     "1",
						UserId: fmt.Sprint(userID),
						Name:   name,
					},
				},
				Total: 1,
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
