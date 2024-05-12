package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/goravel/framework/database/orm"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mocksservice "market.goravel.dev/package/app/mocks/services"
	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
	utilsresponse "market.goravel.dev/utils/response"
)

type PackageControllerSuite struct {
	suite.Suite
	ctx                context.Context
	packageController  *PackageController
	mockLang           *mockstranslation.Translator
	mockPackageService *mocksservice.Package
	mockTagService     *mocksservice.Tag
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
	s.packageController = &PackageController{
		packageService: s.mockPackageService,
		tagService:     s.mockTagService,
	}
}

func (s *PackageControllerSuite) TestCreatePackage() {
	var (
		name   = "goravel"
		url    = "https://goravel.dev"
		userID = uint64(1)

		pkg = models.Package{
			UUIDModel: models.UUIDModel{
				ID: 1,
			},
			Name:   name,
			UserID: userID,
			Link:   url,
		}

		tags = []*models.Tag{
			{
				UUIDModel: models.UUIDModel{
					ID: 1,
				},
				Name:   "goravel",
				UserID: userID,
			},
		}
		mockOrm            *mocksorm.Orm
		mockOrmQuery       *mocksorm.Query
		mockOrmAssociation *mocksorm.Association
		mockLang           *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockFactory.Log()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrmAssociation = mockFactory.OrmAssociation()
		mockLang = mockFactory.Lang(s.ctx)
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name             string
		request          *protopackage.CreatePackageRequest
		setup            func()
		expectedResponse *protopackage.CreatePackageResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
			},
			expectedResponse: &protopackage.CreatePackageResponse{
				Status:  utilsresponse.NewOkStatus(),
				Package: pkg.ToProto(),
			},
		},
		{
			name: "Sad path - Create returns error",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Sad path - Create package with tags",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name = ?", "goravel").Return(mockOrmQuery).Once()
				mockOrmQuery.On("FirstOrFail", mock.AnythingOfType("*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*models.Tag)
						*tagPtr = models.Tag{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}
					}).Once()

				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*models.Package)
						packagePtr.Tags = tags
					}).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*models.Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(nil).Once()
			},
			expectedResponse: &protopackage.CreatePackageResponse{
				Status: utilsresponse.NewOkStatus(),
				Package: &protopackage.Package{
					Id:     "1",
					UserId: fmt.Sprint(userID),
					Name:   name,
					Link:   url,
					Tags: []*protopackage.Tag{
						tags[0].ToProto(),
					},
				},
			},
		},
		{
			name: "Sad path - Tags association error",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name = ?", "goravel").Return(mockOrmQuery).Once()
				mockOrmQuery.On("FirstOrFail", mock.AnythingOfType("*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*models.Tag)
						*tagPtr = models.Tag{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}
					}).Once()
				mockOrmQuery.On("Create", mock.MatchedBy(func(tag *models.Tag) bool {
					return tag.ID > 0 && tag.Name == "goravel" && tag.UserID == userID
				})).Return(nil).Once()

				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*models.Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Sad path - Find tag error",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name = ?", "goravel").Return(mockOrmQuery).Once()
				mockOrmQuery.On("FirstOrFail", mock.AnythingOfType("*models.Tag")).
					Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Happy path - Tags does not exist",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Twice()
				mockOrmQuery.On("Where", "name = ?", "goravel").Return(mockOrmQuery).Once()
				mockOrmQuery.On("FirstOrFail", mock.AnythingOfType("*models.Tag")).
					Return(orm.ErrRecordNotFound).Once()
				mockOrmQuery.On("Create", mock.MatchedBy(func(tag *models.Tag) bool {
					return tag.ID > 0 && tag.Name == "goravel" && tag.UserID == userID
				})).Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*models.Tag)
						*tagPtr = models.Tag{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}
					}).Once()

				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*models.Package)
						packagePtr.Tags = tags
					}).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*models.Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(nil).Once()
			},
			expectedResponse: &protopackage.CreatePackageResponse{
				Status: utilsresponse.NewOkStatus(),
				Package: &protopackage.Package{
					Id:     "1",
					UserId: fmt.Sprint(userID),
					Name:   name,
					Link:   url,
					Tags: []*protopackage.Tag{
						tags[0].ToProto(),
					},
				},
			},
		},
		{
			name: "Sad path - Create Tag error",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   name,
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockOrmQuery.On("Create", mock.MatchedBy(func(pkg *models.Package) bool {
					if pkg.ID == 0 {
						return false
					}
					pkg.ID = 1
					return pkg.Name == name && pkg.UserID == userID && pkg.Link == url
				})).Return(nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Twice()
				mockOrmQuery.On("Where", "name = ?", "goravel").Return(mockOrmQuery).Once()
				mockOrmQuery.On("FirstOrFail", mock.AnythingOfType("*models.Tag")).
					Return(orm.ErrRecordNotFound).Once()
				mockOrmQuery.On("Create", mock.MatchedBy(func(tag *models.Tag) bool {
					return tag.ID > 0 && tag.Name == "goravel" && tag.UserID == userID
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Sad path - Request validation error",
			request: &protopackage.CreatePackageRequest{
				UserId: fmt.Sprint(userID),
				Name:   "",
				Url:    url,
				Tags:   []string{"goravel"},
			},
			setup: func() {
				beforeEach()
				mockLang.On("Get", "required.name").Return("Name is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("Name is required"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.packageController.CreatePackage(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockPackageService.AssertExpectations(s.T())
			s.mockTagService.AssertExpectations(s.T())
		})
	}
}

func (s *PackageControllerSuite) TestGetPackage() {
	var (
		packageID = "1"
		user      = &protouser.User{
			Id:   "1",
			Name: "test",
		}
		tags = []*models.Tag{
			{
				UUIDModel: models.UUIDModel{
					ID: 1,
				},
				Name: "test",
			},
		}
		userID = uint64(1)
		name   = "goravel"
	)

	tests := []struct {
		name             string
		request          *protopackage.GetPackageRequest
		setup            func()
		expectedResponse *protopackage.GetPackageResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.GetPackageRequest{
				Id: packageID,
			},
			setup: func() {
				s.mockPackageService.On("GetPackageByID", packageID).Return(&models.Package{
					UUIDModel: models.UUIDModel{
						ID: 1,
					},
					UserID: userID,
					Name:   name,
					User:   user,
					Tags:   tags,
				}, nil).Once()
			},
			expectedResponse: &protopackage.GetPackageResponse{
				Status: utilsresponse.NewOkStatus(),
				Package: &protopackage.Package{
					Id:     packageID,
					UserId: fmt.Sprint(userID),
					Name:   name,
					User:   user,
					Tags: []*protopackage.Tag{
						{
							Id:   "1",
							Name: "test",
						},
					},
				},
			},
		},
		{
			name:    "Sad path - PackageID is empty",
			request: &protopackage.GetPackageRequest{},
			setup: func() {
				s.mockLang.On("Get", "required.package_id").Return("PackageID is required").Once()
			},
			expectedErr: utilserrors.NewBadRequest("PackageID is required"),
		},
		{
			name: "Sad path - GetPackageByID returns error",
			request: &protopackage.GetPackageRequest{
				Id: packageID,
			},
			setup: func() {
				s.mockPackageService.On("GetPackageByID", packageID).Return(&models.Package{}, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - package not found",
			request: &protopackage.GetPackageRequest{
				Id: packageID,
			},
			setup: func() {
				s.mockPackageService.On("GetPackageByID", packageID).Return(&models.Package{}, nil).Once()
				s.mockLang.On("Get", "not_exist.package").Return("Package not found").Once()
			},
			expectedErr: utilserrors.NewNotFound("Package not found"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.packageController.GetPackage(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockPackageService.AssertExpectations(s.T())
			s.mockTagService.AssertExpectations(s.T())
		})

	}
}

func (s *PackageControllerSuite) TestGetPackages() {
	var (
		users = []*protouser.User{
			{
				Id:   "1",
				Name: "test",
			},
		}
		total      = int64(1)
		userID     = uint64(1)
		name       = "goravel"
		pagination = &protobase.Pagination{
			Page:  1,
			Limit: 10,
		}
		query = &protopackage.PackagesQuery{
			Name: name,
		}
	)

	tests := []struct {
		name             string
		request          *protopackage.GetPackagesRequest
		setup            func()
		expectedResponse *protopackage.GetPackagesResponse
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.GetPackagesRequest{
				Pagination: pagination,
				Query:      query,
			},
			setup: func() {
				total = 1
				s.mockPackageService.On("GetPackages", query, pagination).Return([]*models.Package{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: userID,
						Name:   name,
						User:   users[0],
					},
				}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetPackagesResponse{
				Status: utilsresponse.NewOkStatus(),
				Packages: []*protopackage.Package{
					{
						Id:     "1",
						UserId: fmt.Sprint(userID),
						Name:   name,
						User:   users[0],
						Tags:   []*protopackage.Tag{},
					},
				},
				Total: 1,
			},
		},
		{
			name: "Sad path - GetPackages returns error",
			request: &protopackage.GetPackagesRequest{
				Pagination: pagination,
				Query:      query,
			},
			setup: func() {
				total = 0
				s.mockPackageService.On("GetPackages", query, pagination).Return(nil, total, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Happy path - Packages is empty",
			request: &protopackage.GetPackagesRequest{
				Pagination: pagination,
				Query:      query,
			},
			setup: func() {
				total = 0
				s.mockPackageService.On("GetPackages", query, pagination).Return([]*models.Package{}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetPackagesResponse{
				Status:   utilsresponse.NewOkStatus(),
				Packages: []*protopackage.Package{},
			},
		},
		{
			name: "Happy path - pagination is nil",
			request: &protopackage.GetPackagesRequest{
				Query: query,
			},
			setup: func() {
				total = 0
				s.mockPackageService.On("GetPackages", query, &protobase.Pagination{Page: 1, Limit: 10}).Return([]*models.Package{}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetPackagesResponse{
				Status:   utilsresponse.NewOkStatus(),
				Packages: []*protopackage.Package{},
			},
		},
		{
			name: "Happy path - query is nil",
			request: &protopackage.GetPackagesRequest{
				Pagination: pagination,
			},
			setup: func() {
				total = 1
				s.mockPackageService.On("GetPackages", (*protopackage.PackagesQuery)(nil), pagination).Return([]*models.Package{
					{
						UUIDModel: models.UUIDModel{
							ID: 1,
						},
						UserID: userID,
						Name:   name,
						User:   users[0],
					},
				}, total, nil).Once()
			},
			expectedResponse: &protopackage.GetPackagesResponse{
				Status: utilsresponse.NewOkStatus(),
				Packages: []*protopackage.Package{
					{
						Id:     "1",
						UserId: fmt.Sprint(userID),
						Name:   name,
						User:   users[0],
						Tags:   []*protopackage.Tag{},
					},
				},
				Total: 1,
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.packageController.GetPackages(s.ctx, test.request)
			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)

			s.mockLang.AssertExpectations(s.T())
			s.mockTagService.AssertExpectations(s.T())
		})
	}
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
