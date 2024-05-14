package models

import (
	"errors"
	"testing"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/database/orm"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	"github.com/goravel/framework/support/carbon"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	protopackage "market.goravel.dev/proto/package"
	utilserrors "market.goravel.dev/utils/errors"
)

type PackageSuite struct {
	suite.Suite
	pkg *Package
}

func TestPackageSuite(t *testing.T) {
	suite.Run(t, new(PackageSuite))
}

func (s *PackageSuite) SetupTest() {
	s.pkg = NewPackage()
}

func (s *PackageSuite) TestAttachTags() {
	var (
		tags   = []string{"goravel"}
		name   = "goravel"
		url    = "https://goravel.dev"
		userID = uint64(1)

		pkg = Package{
			UUIDModel: UUIDModel{
				ID: 1,
			},
			Name:   name,
			UserID: userID,
			Link:   url,
		}

		tagModels = []*Tag{
			{
				UUIDModel: UUIDModel{
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

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
		mockFactory.Log()
		mockOrmAssociation = mockFactory.OrmAssociation()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name        string
		setup       func()
		expectedErr error
	}{
		{
			name: "Happy path - Create package with tags",
			setup: func() {
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*[]*Tag)
						*tagPtr = []*Tag{{UUIDModel: UUIDModel{ID: 1}, Name: "goravel", UserID: userID}}
					}).Once()

				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*Package)
						packagePtr.Tags = tagModels
					}).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(nil).Once()
			},
		},
		{
			name: "Sad path - Tags association error",
			setup: func() {
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).(*[]*Tag)
						*tagPtr = []*Tag{{UUIDModel: UUIDModel{ID: 1}, Name: "goravel", UserID: userID}}
					}).Once()

				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*Package)
						packagePtr.Tags = tagModels
					}).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Sad path - Find tag error",
			setup: func() {
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
		{
			name: "Happy path - Tags do not exist",
			setup: func() {
				mockOrm.On("Query").Return(mockOrmQuery).Twice()
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).Once()

				mockOrmQuery.On("Create", mock.MatchedBy(func(tags []*Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel" && tags[0].UserID == userID && tags[0].IsShow == 1
				})).Return(nil).
					Run(func(args mock.Arguments) {
						tagPtr := args.Get(0).([]*Tag)
						tagPtr[0] = &Tag{UUIDModel: UUIDModel{ID: 1}, Name: "goravel", UserID: userID}
					}).Once()

				mockOrmQuery.On("Model", mock.AnythingOfType("*models.Package")).Return(mockOrmQuery).
					Run(func(args mock.Arguments) {
						packagePtr := args.Get(0).(*Package)
						packagePtr.Tags = tagModels
					}).Once()
				mockOrmQuery.On("Association", "Tags").Return(mockOrmAssociation).Once()
				mockOrmAssociation.On("Replace", mock.MatchedBy(func(tags []*Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel"
				})).Return(nil).Once()
			},
		},
		{
			name: "Sad path - Create Tag error",
			setup: func() {
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("WhereIn", "name", []any{"goravel"}).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Find", mock.AnythingOfType("*[]*models.Tag")).
					Return(nil).Once()

				mockOrmQuery.On("Create", mock.MatchedBy(func(tags []*Tag) bool {
					return len(tags) == 1 && tags[0].Name == "goravel" && tags[0].UserID == userID && tags[0].IsShow == 1
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			err := s.pkg.AttachTags(&pkg, tags)
			s.Equal(test.expectedErr, err)

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
			mockOrmAssociation.AssertExpectations(s.T())
		})
	}
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
		mockFactory.Log()
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
				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("First", &pack).Run(func(args mock.Arguments) {
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
			name: "Sad path - get package error",
			setup: func() {
				var pack Package
				mockOrmQuery.On("With", "Tags", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("First", &pack).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			returnedPackage, err := s.pkg.GetPackageByID(id, fields)

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
		UserID:        uint64(userID),
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
		Tags:          []*protopackage.Tag{},
	}, pack.ToProto())
}

func (s *PackageSuite) TestUpdatePackage() {
	var (
		name   = "goravel"
		url    = "https://goravel.dev"
		userID = uint64(1)

		pkg = Package{
			UUIDModel: UUIDModel{
				ID: 1,
			},
			Name:   name,
			UserID: userID,
			Link:   url,
		}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
		mockFactory.Log()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name        string
		setup       func()
		expectedErr error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockOrmQuery.On("Save", mock.MatchedBy(func(pkg *Package) bool {
					return pkg.Name == name && pkg.Link == url && pkg.UserID == userID
				})).Return(nil).Once()
			},
			expectedErr: nil,
		},
		{
			name: "Sad path - save package error",
			setup: func() {
				mockOrmQuery.On("Save", mock.MatchedBy(func(pkg *Package) bool {
					return pkg.Name == name && pkg.Link == url && pkg.UserID == userID
				})).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			err := s.pkg.UpdatePackage(&pkg)

			s.Equal(test.expectedErr, err)

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}
