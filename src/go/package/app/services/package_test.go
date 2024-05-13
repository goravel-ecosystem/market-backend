package services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/goravel/framework/contracts/http"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	"github.com/goravel/framework/support/carbon"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	protopackage "market.goravel.dev/proto/package"

	mocks "market.goravel.dev/package/app/mocks/models"
	mocksservice "market.goravel.dev/package/app/mocks/services"
	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

type PackageTestSuite struct {
	suite.Suite
	ctx                  context.Context
	packageImpl          *PackageImpl
	mockPackageInterface *mocks.PackageInterface
	mockUserService      *mocksservice.User
	mockLang             *mockstranslation.Translator
}

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}

func (s *PackageTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockUserService = &mocksservice.User{}
	s.mockPackageInterface = &mocks.PackageInterface{}
	mockFactory := testingmock.Factory()
	s.mockLang = mockFactory.Lang(s.ctx)
	s.packageImpl = &PackageImpl{
		packageModel: s.mockPackageInterface,
		userService:  s.mockUserService,
	}
}

func (s *PackageTestSuite) TestCreatePackage() {
	var (
		name          = "goravel"
		url           = "https://goravel.dev"
		userID        = uint64(1)
		lastUpdatedAt = carbon.Now().String()

		pkg = models.Package{
			UUIDModel: models.UUIDModel{
				ID: 1,
			},
			Name:          name,
			UserID:        userID,
			Link:          url,
			LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(lastUpdatedAt)},
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
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockFactory.Log()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrmAssociation = mockFactory.OrmAssociation()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name             string
		request          *protopackage.CreatePackageRequest
		setup            func()
		expectedResponse *models.Package
		expectedErr      error
	}{
		{
			name: "Happy path",
			request: &protopackage.CreatePackageRequest{
				UserId:        fmt.Sprint(userID),
				Name:          name,
				Url:           url,
				LastUpdatedAt: lastUpdatedAt,
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
			expectedResponse: &pkg,
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
			name: "Happy path - Create package with tags",
			request: &protopackage.CreatePackageRequest{
				UserId:        fmt.Sprint(userID),
				Name:          name,
				Url:           url,
				LastUpdatedAt: lastUpdatedAt,
				Tags:          []string{"goravel"},
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
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*[]*models.Tag)
						*tagPtr = []*models.Tag{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}}
					}).Once()

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
			expectedResponse: &models.Package{
				UUIDModel: models.UUIDModel{
					ID: 1,
				},
				Name:          name,
				UserID:        userID,
				Link:          url,
				LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(lastUpdatedAt)},
				Tags:          tags,
			},
		},
		{
			name: "Sad path - Tags association error",
			request: &protopackage.CreatePackageRequest{
				UserId:        fmt.Sprint(userID),
				Name:          name,
				Url:           url,
				LastUpdatedAt: lastUpdatedAt,
				Tags:          []string{"goravel"},
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
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*[]*models.Tag)
						*tagPtr = []*models.Tag{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}}
					}).Once()

				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*models.Package)
						packagePtr.Tags = tags
					}).Once()
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
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Happy path - Tags do not exist",
			request: &protopackage.CreatePackageRequest{
				UserId:        fmt.Sprint(userID),
				Name:          name,
				Url:           url,
				LastUpdatedAt: lastUpdatedAt,
				Tags:          []string{"goravel"},
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
				mockOrm.On("Query").Return(mockOrmQuery).Times(3)
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).Once()

				mockOrmQuery.On("Create", mock.MatchedBy(func(tags []*models.Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel" && tags[0].UserID == userID && tags[0].IsShow == 1
				})).Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).([]*models.Tag)
						tagPtr[0] = &models.Tag{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel", UserID: userID}
					}).Once()

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
			expectedResponse: &models.Package{
				UUIDModel: models.UUIDModel{
					ID: 1,
				},
				Name:          name,
				UserID:        userID,
				Link:          url,
				LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(lastUpdatedAt)},
				Tags:          tags,
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
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).Once()

				mockOrmQuery.On("Create", mock.MatchedBy(func(tags []*models.Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel" && tags[0].UserID == userID && tags[0].IsShow == 1
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()
			response, err := s.packageImpl.CreatePackage(test.request)

			s.Equal(test.expectedResponse, response)
			s.Equal(test.expectedErr, err)
		})
	}
}

func (s *PackageTestSuite) TestGetPackageByID() {
	var (
		packageID = "1"
		userID    = uint64(1)
		fields    = []string{"id", "name", "user_id", "summary", "description", "link", "version", "last_updated_at", "view_count"}
		user      = &protouser.User{
			Id:   "1",
			Name: "test",
		}
	)

	tests := []struct {
		name          string
		setup         func()
		expectPackage *models.Package
		expectedErr   error
	}{
		{
			name: "Happy path - GetPackageByID with ID",
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, fields).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockUserService.On("GetUser", s.ctx, userID).Return(user, nil).Once()
			},
			expectPackage: &models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID, User: user},
			expectedErr:   nil,
		},
		{
			name: "Sad path - GetPackageByID returns error",
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, fields).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - GetUser returns error",
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, fields).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockUserService.On("GetUser", s.ctx, userID).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()

			pkg, err := s.packageImpl.GetPackageByID(packageID)
			if test.expectedErr != nil {
				s.Nil(pkg)
				s.Equal(test.expectedErr, err)
			} else {
				s.Nil(err)
				s.Equal(test.expectPackage, pkg)
			}
		})
	}
}

func (s *PackageTestSuite) TestGetPackages() {
	var (
		name   = "go"
		userID = uint64(1)
		fields = []string{"id", "name", "user_id", "summary", "link", "view_count"}
		users  = []*protouser.User{
			{
				Id:   "1",
				Name: "test",
			},
		}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		pagination   *protobase.Pagination
		query        *protopackage.PackagesQuery
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockFactory.Log()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name           string
		setup          func()
		expectPackages []*models.Package
		expectedTotal  int64
		expectedErr    error
	}{
		{
			name: "Happy path - GetPackages with query and pagination",
			setup: func() {
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}
				query = &protopackage.PackagesQuery{
					Name:     name,
					Category: "hot",
				}

				beforeSetup()
				mockOrmQuery.On("OrderByDesc", "view_count").Return(mockOrmQuery).Once()
				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Package"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						packagesPtr := args.Get(2).(*[]*models.Package)
						*packagesPtr = []*models.Package{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 1
					}).Once()
				s.mockUserService.On("GetUsers", s.ctx, []string{fmt.Sprint(userID)}).Return(users, nil).Once()
			},
			expectPackages: []*models.Package{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID, User: users[0]}},
			expectedTotal:  1,
			expectedErr:    nil,
		},
		{
			name: "Happy path - GetPackages without query",
			setup: func() {
				name = ""
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				query = &protopackage.PackagesQuery{}

				beforeSetup()

				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Package"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						packagesPtr := args.Get(2).(*[]*models.Package)
						*packagesPtr = []*models.Package{{UUIDModel: models.UUIDModel{ID: 3}, Name: "goravel/cloudinary", UserID: userID}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 1
					}).Once()
				s.mockUserService.On("GetUsers", s.ctx, []string{fmt.Sprint(userID)}).Return(users, nil).Once()
			},
			expectPackages: []*models.Package{{UUIDModel: models.UUIDModel{ID: 3}, Name: "goravel/cloudinary", UserID: userID, User: users[0]}},
			expectedTotal:  1,
			expectedErr:    nil,
		},
		{
			name: "Sad path - Paginate return error",
			setup: func() {
				name = "go"
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				query = &protopackage.PackagesQuery{
					Name: name,
				}

				beforeSetup()
				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Package"), mock.AnythingOfType("*int64")).
					Return(errors.New("paginate error")).Once()
			},
			expectPackages: nil,
			expectedTotal:  0,
			expectedErr:    utilserrors.New(http.StatusInternalServerError, "paginate error"),
		},
		{
			name: "Happy path - GetUsers returns error",
			setup: func() {
				name = ""
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				query = &protopackage.PackagesQuery{}

				beforeSetup()

				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Package"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						packagesPtr := args.Get(2).(*[]*models.Package)
						*packagesPtr = []*models.Package{{UUIDModel: models.UUIDModel{ID: 3}, Name: "goravel/cloudinary", UserID: userID}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 1
					}).Once()
				s.mockUserService.On("GetUsers", s.ctx, []string{fmt.Sprint(userID)}).Return(nil, errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()

			packages, total, err := s.packageImpl.GetPackages(query, pagination)
			if test.expectedErr != nil {
				s.Nil(packages)
				s.Equal(int64(0), total)
				s.Equal(test.expectedErr, err)
			} else {
				s.Nil(err)
				s.Equal(test.expectPackages, packages)
				s.Equal(test.expectedTotal, total)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *PackageTestSuite) TestUpdatePackage() {
	var (
		packageID     = "1"
		userID        = uint64(1)
		name          = "goravel/gin"
		url           = "https://github.com/goravel/gin"
		lastUpdatedAt = carbon.Now().String()
		tags          = []string{"goravel"}
	)

	tests := []struct {
		name          string
		request       *protopackage.UpdatePackageRequest
		setup         func()
		expectPackage *models.Package
		expectedErr   error
	}{
		{
			name: "Happy path - UpdatePackage with ID",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockPackageInterface.On("UpdatePackage", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				})).Return(nil).Once()
			},
			expectPackage: &models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: name, UserID: userID, Link: url, LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(lastUpdatedAt)}},
		},
		{
			name: "Sad path - GetPackageByID returns error",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(nil, errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - Package does not exist",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{}, nil).Once()
				s.mockLang.On("Get", "not_exist.package").Return("package not exist").Once()
			},
			expectedErr: utilserrors.New(http.StatusNotFound, "package not exist"),
		},
		{
			name: "Sad path - User isn't owner of package",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: 2}, nil).Once()
				s.mockLang.On("Get", "forbidden.update_package").Return("forbidden.update_package").Once()
			},
			expectedErr: utilserrors.New(http.StatusUnauthorized, "forbidden.update_package"),
		},
		{
			name: "Happy path - UpdatePackage with tags",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
				Tags:          tags,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockPackageInterface.On("UpdatePackage", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				})).Return(nil).Once()
				s.mockPackageInterface.On("AttachTags", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				}), tags).Return(nil).Once()
			},
			expectPackage: &models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: name, UserID: userID, Link: url, LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(lastUpdatedAt)}},
		},
		{
			name: "Sad path - UpdatePackage returns error",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
				Tags:          tags,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockPackageInterface.On("UpdatePackage", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				})).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - AttachTags returns error",
			request: &protopackage.UpdatePackageRequest{
				Id:            packageID,
				Name:          name,
				Url:           url,
				UserId:        fmt.Sprint(userID),
				LastUpdatedAt: lastUpdatedAt,
				Tags:          tags,
			},
			setup: func() {
				s.mockPackageInterface.On("GetPackageByID", packageID, []string{}).Return(&models.Package{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin", UserID: userID}, nil).Once()
				s.mockPackageInterface.On("UpdatePackage", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				})).Return(nil).Once()
				s.mockPackageInterface.On("AttachTags", mock.MatchedBy(func(pkg *models.Package) bool {
					return pkg.Name == name && pkg.Link == url
				}), tags).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()

			pkg, err := s.packageImpl.UpdatePackage(s.ctx, test.request)
			if test.expectedErr != nil {
				s.Nil(pkg)
				s.Equal(test.expectedErr, err)
			} else {
				s.Nil(err)
				s.Equal(test.expectPackage, pkg)
			}
		})
	}
}
