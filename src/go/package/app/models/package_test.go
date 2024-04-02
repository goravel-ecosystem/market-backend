package models

import (
	"errors"
	"testing"

	"github.com/goravel/framework/database/orm"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	"github.com/goravel/framework/support/carbon"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	protopackage "market.goravel.dev/proto/package"
)

type PackageSuite struct {
	suite.Suite
	pack *Package
}

func TestPackageSuite(t *testing.T) {
	suite.Run(t, new(PackageSuite))
}

func (s *PackageSuite) SetupTest() {
	s.pack = NewPackage()
}

func (s *PackageSuite) TestGetPackageByID() {
	var (
		id     = "1"
		fields = []string{"name"}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		pack         Package
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
		mockOrmQuery.On("Where", "id", id).Return(mockOrmQuery).Once()
		mockOrmQuery.On("Select", []string{"name"}).Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name          string
		setup         func()
		expectPackage *Package
		expectedErr   error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockOrmQuery.On("FirstOrFail", &pack).Run(func(args mock.Arguments) {
					pack := args.Get(0).(*Package)
					pack.ID = 1
					pack.Name = "Goravel"
				}).Return(nil).Once()
			},
			expectPackage: &Package{
				UUIDModel: UUIDModel{
					ID: 1,
				},
				Name: "Goravel",
			},
		},
		{
			name: "Sad path - get user error",
			setup: func() {
				var pack Package
				mockOrmQuery.On("FirstOrFail", &pack).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			returnedPackage, err := s.pack.GetPackageByID(id, fields)

			if test.expectedErr != nil {
				s.Nil(returnedPackage)
				s.Equal(test.expectedErr, err)
			} else {
				s.NotNil(returnedPackage)
				s.Nil(err)
				s.Equal(test.expectPackage, returnedPackage)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *PackageSuite) TestToProto() {
	var (
		id            = 1
		userID        = 1
		name          = "Goravel"
		summary       = "summary"
		description   = "description"
		link          = "https://goravel.dev"
		version       = "v1.0.0"
		lastUpdatedAt = carbon.DateTime{}
		createAt      = carbon.DateTime{}
		updatedAt     = carbon.DateTime{}
	)

	pack := Package{
		UUIDModel: UUIDModel{
			ID: uint64(id),
			Timestamps: orm.Timestamps{
				CreatedAt: createAt,
				UpdatedAt: updatedAt,
			},
		},
		UserID:        uint(userID),
		Name:          name,
		Summary:       summary,
		Description:   description,
		Link:          link,
		Version:       version,
		LastUpdatedAt: lastUpdatedAt,
	}

	s.Equal(&protopackage.Package{
		Id:            "1",
		UserId:        "1",
		Name:          name,
		Summary:       summary,
		Description:   description,
		Link:          link,
		Version:       version,
		LastUpdatedAt: lastUpdatedAt.ToString(),
		CreatedAt:     createAt.ToString(),
		UpdatedAt:     updatedAt.ToString(),
	}, pack.ToProto())
}