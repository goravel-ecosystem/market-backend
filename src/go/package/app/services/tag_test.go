package services

import (
	"errors"
	"testing"

	mocksorm "github.com/goravel/framework/mocks/database/orm"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
)

type TagTestSuite struct {
	suite.Suite
	tagImpl *TagImpl
}

func TestTagTestSuite(t *testing.T) {
	suite.Run(t, new(TagTestSuite))
}

func (s *TagTestSuite) SetupTest() {
	s.tagImpl = NewTagImpl()
}

func (s *TagTestSuite) TestGetTags() {
	var (
		packageID = "1"
		name      = "go"
		fields    = []string{"id", "name"}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		pagination   *protobase.Pagination
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name          string
		setup         func()
		expectTags    []*models.Tag
		expectedTotal int64
		expectedErr   error
	}{
		{
			name: "Happy path - GetTags with packageID and name",
			setup: func() {
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				beforeSetup()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Table", "package_tags").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "package_id = ?", packageID).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Pluck", "tag_id", mock.Anything).Return(nil).Once()
				mockOrmQuery.On("WhereIn", "id", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Tag"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagsPtr := args.Get(2).(*[]*models.Tag)
						*tagsPtr = []*models.Tag{{UUIDModel: models.UUIDModel{ID: 1}, Name: "GoLang"}, {UUIDModel: models.UUIDModel{ID: 2}, Name: "Gopher"}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 2
					}).Once()
			},
			expectTags:    []*models.Tag{{UUIDModel: models.UUIDModel{ID: 1}, Name: "GoLang"}, {UUIDModel: models.UUIDModel{ID: 2}, Name: "Gopher"}},
			expectedTotal: 2,
			expectedErr:   nil,
		},
		{
			name: "Happy path - GetTags with only packageID",
			setup: func() {
				packageID = "1"
				name = ""
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				beforeSetup()

				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Table", "package_tags").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "package_id = ?", packageID).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Pluck", "tag_id", mock.Anything).Return(nil).Once()
				mockOrmQuery.On("WhereIn", "id", mock.Anything).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Tag"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagsPtr := args.Get(2).(*[]*models.Tag)
						*tagsPtr = []*models.Tag{{UUIDModel: models.UUIDModel{ID: 3}, Name: "Python"}, {UUIDModel: models.UUIDModel{ID: 4}, Name: "Java"}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 2
					}).Once()
			},
			expectTags:    []*models.Tag{{UUIDModel: models.UUIDModel{ID: 3}, Name: "Python"}, {UUIDModel: models.UUIDModel{ID: 4}, Name: "Java"}},
			expectedTotal: 2,
			expectedErr:   nil,
		},
		{
			name: "Happy path - GetTags with only name",
			setup: func() {
				packageID = ""
				name = "go"
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				beforeSetup()

				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Tag"), mock.AnythingOfType("*int64")).
					Return(nil).
					Run(func(args mock.Arguments) {
						tagsPtr := args.Get(2).(*[]*models.Tag)
						*tagsPtr = []*models.Tag{{UUIDModel: models.UUIDModel{ID: 5}, Name: "GoTest"}, {UUIDModel: models.UUIDModel{ID: 6}, Name: "JavaScript"}}

						totalPtr := args.Get(3).(*int64)
						*totalPtr = 2
					}).Once()
			},
			expectTags:    []*models.Tag{{UUIDModel: models.UUIDModel{ID: 5}, Name: "GoTest"}, {UUIDModel: models.UUIDModel{ID: 6}, Name: "JavaScript"}},
			expectedTotal: 2,
			expectedErr:   nil,
		},
		{
			name: "Sad path - Pluck return error",
			setup: func() {
				packageID = "1"

				beforeSetup()

				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Table", "package_tags").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Where", "package_id = ?", packageID).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Pluck", "tag_id", mock.Anything).Return(errors.New("pluck error")).Once()
			},
			expectTags:    nil,
			expectedTotal: 0,
			expectedErr:   errors.New("pluck error"),
		},
		{
			name: "Sad path - Paginate return error",
			setup: func() {
				packageID = ""
				name = "go"
				pagination = &protobase.Pagination{
					Page:  1,
					Limit: 10,
				}

				beforeSetup()

				mockOrmQuery.On("Where", "name LIKE ?", "%"+name+"%").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Select", fields).Return(mockOrmQuery).Once()
				mockOrmQuery.On("Paginate", int(pagination.GetPage()), int(pagination.GetLimit()), mock.AnythingOfType("*[]*models.Tag"), mock.AnythingOfType("*int64")).
					Return(errors.New("paginate error")).Once()
			},
			expectTags:    nil,
			expectedTotal: 0,
			expectedErr:   errors.New("paginate error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			test.setup()

			tags, total, err := s.tagImpl.GetTags(packageID, name, pagination)
			if test.expectedErr != nil {
				s.Nil(tags)
				s.Equal(int64(0), total)
				s.Equal(test.expectedErr, err)
			} else {
				s.Nil(err)
				s.Equal(test.expectTags, tags)
				s.Equal(test.expectedTotal, total)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}
