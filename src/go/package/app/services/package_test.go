package services

import (
	"errors"
	"testing"

	"github.com/goravel/framework/contracts/http"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	protopackage "market.goravel.dev/proto/package"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	utilserrors "market.goravel.dev/utils/errors"
)

type PackageTestSuite struct {
	suite.Suite
	packageImpl *PackageImpl
}

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(PackageTestSuite))
}

func (s *PackageTestSuite) SetupTest() {
	s.packageImpl = NewPackageImpl()
}

func (s *PackageTestSuite) TestGetTags() {
	var (
		name   = "go"
		fields = []string{"id", "name", "user_id", "summary", "link", "view_count"}

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
				mockOrmQuery.On("OrderBy", "view_count").Return(mockOrmQuery).Once()
				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Package"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						packagesPtr := args.Get(2).(*[]*models.Package)
						*packagesPtr = []*models.Package{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin"}, {UUIDModel: models.UUIDModel{ID: 2}, Name: "goravel/fiber"}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 2
					}).Once()
			},
			expectPackages: []*models.Package{{UUIDModel: models.UUIDModel{ID: 1}, Name: "goravel/gin"}, {UUIDModel: models.UUIDModel{ID: 2}, Name: "goravel/fiber"}},
			expectedTotal:  2,
			expectedErr:    nil,
		},
		{
			name: "Happy path - GetTags without query",
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
						*packagesPtr = []*models.Package{{UUIDModel: models.UUIDModel{ID: 3}, Name: "goravel/cloudinary"}, {UUIDModel: models.UUIDModel{ID: 4}, Name: "goravel/minio"}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 2
					}).Once()
			},
			expectPackages: []*models.Package{{UUIDModel: models.UUIDModel{ID: 3}, Name: "goravel/cloudinary"}, {UUIDModel: models.UUIDModel{ID: 4}, Name: "goravel/minio"}},
			expectedTotal:  2,
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
