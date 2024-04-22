package models

import (
	"errors"
	"net/http"
	"testing"

	mocksorm "github.com/goravel/framework/mocks/database/orm"
	mockshash "github.com/goravel/framework/mocks/hash"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
)

type UserSuite struct {
	suite.Suite
	user *User
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) SetupTest() {
	s.user = NewUser()
}

func (s *UserSuite) TestGetUserByEmail() {
	var (
		email  = "hello@goravel.dev"
		fields = []string{"id"}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		user         User
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
		mockOrmQuery.On("Where", "email", email).Return(mockOrmQuery).Once()
		mockOrmQuery.On("Select", []string{"id"}).Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name        string
		setup       func()
		expectUser  *User
		expectedErr error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockOrmQuery.On("First", &user).Run(func(args mock.Arguments) {
					user := args.Get(0).(*User)
					user.ID = 1
				}).Return(nil).Once()
			},
		},
		{
			name: "Sad path - get user error",
			setup: func() {
				var user User
				mockOrmQuery.On("First", &user).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			returnedUser, err := s.user.GetUserByEmail(email, fields)

			if test.expectedErr != nil {
				s.Nil(returnedUser)
				s.Equal(test.expectedErr, err)
			} else {
				s.NotNil(returnedUser)
				s.Nil(err)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *UserSuite) TestGetUserByID() {
	var (
		id     = "1"
		fields = []string{"name"}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		user         User
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockFactory.Log()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
		mockOrmQuery.On("Where", "id", id).Return(mockOrmQuery).Once()
		mockOrmQuery.On("Select", []string{"name"}).Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name        string
		setup       func()
		expectUser  *User
		expectedErr error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockOrmQuery.On("First", &user).Run(func(args mock.Arguments) {
					user := args.Get(0).(*User)
					user.ID = 1
					user.Name = "Goravel"
				}).Return(nil).Once()
			},
		},
		{
			name: "Sad path - get user error",
			setup: func() {
				var user User
				mockOrmQuery.On("First", &user).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			returnedUser, err := s.user.GetUserByID(id, fields)

			if test.expectedErr != nil {
				s.Nil(returnedUser)
				s.Equal(test.expectedErr, err)
			} else {
				s.NotNil(returnedUser)
				s.Nil(err)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *UserSuite) TestGetUsers() {
	var (
		ids    = []string{"1"}
		fields = []string{"name"}

		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
		users        []*User
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockOrm = mockFactory.Orm()
		mockFactory.Log()
		mockOrmQuery = mockFactory.OrmQuery()
		mockOrm.On("Query").Return(mockOrmQuery).Once()
		mockOrmQuery.On("WhereIn", "id", []any{"1"}).Return(mockOrmQuery).Once()
		mockOrmQuery.On("Select", []string{"name"}).Return(mockOrmQuery).Once()
	}

	tests := []struct {
		name        string
		setup       func()
		expectUsers []*User
		expectedErr error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockOrmQuery.On("Find", &users).Run(func(args mock.Arguments) {
					users := args.Get(0).(*[]*User)
					*users = []*User{
						{
							UUIDModel: UUIDModel{
								ID: 1,
							},
							Name: "Krishna",
						},
					}
				}).Return(nil).Once()
			},
			expectUsers: []*User{
				{
					UUIDModel: UUIDModel{
						ID: 1,
					},
					Name: "Krishna",
				},
			},
		},
		{
			name: "Sad path - get users error",
			setup: func() {
				var users []*User
				mockOrmQuery.On("Find", &users).Return(errors.New("error")).Once()
			},
			expectedErr: utilserrors.New(http.StatusInternalServerError, "error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			returnedUsers, err := s.user.GetUsers(ids, fields)

			if test.expectedErr != nil {
				s.Nil(returnedUsers)
				s.Equal(test.expectedErr, err)
			} else {
				s.NotNil(returnedUsers)
				s.Nil(err)
			}

			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *UserSuite) TestRegister() {
	var (
		email          = "hello@goravel.dev"
		name           = "Goravel"
		password       = "password"
		hashedPassword = "hashed_password"

		mockHash     *mockshash.Hash
		mockOrm      *mocksorm.Orm
		mockOrmQuery *mocksorm.Query
	)

	beforeSetup := func() {
		mockFactory := testingmock.Factory()
		mockHash = mockFactory.Hash()
		mockOrm = mockFactory.Orm()
		mockOrmQuery = mockFactory.OrmQuery()
	}

	tests := []struct {
		name        string
		setup       func()
		expectUser  *User
		expectedErr error
	}{
		{
			name: "Happy path",
			setup: func() {
				mockHash.On("Make", password).Return(hashedPassword, nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Create", mock.MatchedBy(func(user *User) bool {
					return user.ID > 0 && user.Name == name && user.Email == email && user.Password == hashedPassword
				})).Return(nil).Once()
			},
		},
		{
			name: "Sad path - make hash password error",
			setup: func() {
				mockHash.On("Make", password).Return("", errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - create user error",
			setup: func() {
				mockHash.On("Make", password).Return(hashedPassword, nil).Once()
				mockOrm.On("Query").Return(mockOrmQuery).Once()
				mockOrmQuery.On("Create", mock.MatchedBy(func(user *User) bool {
					return user.ID > 0 && user.Name == name && user.Email == email && user.Password == hashedPassword
				})).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeSetup()
			test.setup()
			user, err := s.user.Register(name, email, password)

			if test.expectedErr != nil {
				s.Nil(user)
				s.Equal(test.expectedErr, err)
			} else {
				s.NotNil(user)
				s.Nil(err)
			}

			mockHash.AssertExpectations(s.T())
			mockOrm.AssertExpectations(s.T())
			mockOrmQuery.AssertExpectations(s.T())
		})
	}
}

func (s *UserSuite) TestToProto() {
	var (
		id      = 1
		name    = "Goravel"
		email   = "hello@goravel.dev"
		avatar  = "avatar"
		summary = "summary"
	)

	user := User{
		UUIDModel: UUIDModel{
			ID: uint64(id),
		},
		Name:    name,
		Email:   email,
		Avatar:  avatar,
		Summary: summary,
	}

	s.Equal(&protouser.User{
		Id:      "1",
		Name:    name,
		Email:   email,
		Avatar:  avatar,
		Summary: summary,
	}, user.ToProto())
}
