package services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/goravel/framework/contracts/http"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
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
}

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}

func (s *PackageTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockUserService = &mocksservice.User{}
	s.mockPackageInterface = &mocks.PackageInterface{}
	s.packageImpl = &PackageImpl{
		packageModel: s.mockPackageInterface,
		userService:  s.mockUserService,
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
