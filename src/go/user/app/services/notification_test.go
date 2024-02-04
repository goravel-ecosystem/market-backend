package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/contracts/translation"
	mockscache "github.com/goravel/framework/mocks/cache"
	mocksconfig "github.com/goravel/framework/mocks/config"
	mocksmail "github.com/goravel/framework/mocks/mail"
	mockstranslation "github.com/goravel/framework/mocks/translation"
	testingmock "github.com/goravel/framework/testing/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	notificationImpl *NotificationImpl
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupTest() {
	s.notificationImpl = NewNotificationImpl()
}

func (s *AuthTestSuite) TestSendEmailRegisterCode() {
	var (
		ctx        = context.Background()
		email      = "hello@goravel.dev"
		mockCache  *mockscache.Cache
		mockConfig *mocksconfig.Config
		mockMail   *mocksmail.Mail
		mockLang   *mockstranslation.Translator
	)

	beforeEach := func() {
		mockFactory := testingmock.Factory()
		mockCache = mockFactory.Cache()
		mockConfig = mockFactory.Config()
		mockMail = mockFactory.Mail()
		mockLang = mockFactory.Lang(ctx)
	}

	tests := []struct {
		name           string
		setup          func()
		expectedKeyLen int
		expectedErr    error
	}{
		{
			name: "Happy path - running in production",
			setup: func() {
				mockConfig.On("GetString", "app.env").Return("production").Twice()
				mockCache.On("Put", mock.MatchedBy(func(key string) bool {
					return len(key) == 32
				}), mock.MatchedBy(func(key string) bool {
					return len(key) == 6
				}), 300*time.Second).Return(nil).Once()
				mockLang.On("Get", "register_code.subject", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("subject").Once()
				mockLang.On("Get", "register_code.content", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("html").Once()
				mockMail.On("To", []string{email}).Return(mockMail).Once()
				mockMail.On("Content", mail.Content{Subject: "subject", Html: "html"}).Return(mockMail).Once()
				mockMail.On("Queue").Return(nil).Once()
			},
			expectedKeyLen: 32,
		},
		{
			name: "Happy path - running in development",
			setup: func() {
				mockConfig.On("GetString", "app.env").Return("development").Times(4)
				mockCache.On("Put", mock.MatchedBy(func(key string) bool {
					return len(key) == 32
				}), mock.MatchedBy(func(key string) bool {
					return len(key) == 6
				}), 300*time.Second).Return(nil).Once()
				mockLang.On("Get", "register_code.subject", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("subject").Once()
				mockLang.On("Get", "register_code.content", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("html").Once()
				mockMail.On("To", []string{email}).Return(mockMail).Once()
				mockMail.On("Content", mail.Content{Subject: "subject", Html: "html"}).Return(mockMail).Once()
				mockMail.On("Queue").Return(nil).Once()
			},
			expectedKeyLen: 32,
		},
		{
			name: "Happy path - running in local",
			setup: func() {
				mockConfig.On("GetString", "app.env").Return("local").Times(4)
				mockCache.On("Put", mock.MatchedBy(func(key string) bool {
					return len(key) == 32
				}), mock.MatchedBy(func(key string) bool {
					return len(key) == 6
				}), 300*time.Second).Return(nil).Once()
			},
			expectedKeyLen: 32,
		},
		{
			name: "Sad path - send email failed",
			setup: func() {
				mockConfig.On("GetString", "app.env").Return("production").Twice()
				mockCache.On("Put", mock.MatchedBy(func(key string) bool {
					return len(key) == 32
				}), mock.MatchedBy(func(key string) bool {
					return len(key) == 6
				}), 300*time.Second).Return(nil).Once()
				mockLang.On("Get", "register_code.subject", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("subject").Once()
				mockLang.On("Get", "register_code.content", mock.MatchedBy(func(option translation.Option) bool {
					return len(option.Replace["code"]) == 6
				})).Return("html").Once()
				mockMail.On("To", []string{email}).Return(mockMail).Once()
				mockMail.On("Content", mail.Content{Subject: "subject", Html: "html"}).Return(mockMail).Once()
				mockMail.On("Queue").Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Sad path - put cache failed",
			setup: func() {
				mockConfig.On("GetString", "app.env").Return("production").Once()
				mockCache.On("Put", mock.MatchedBy(func(key string) bool {
					return len(key) == 32
				}), mock.MatchedBy(func(key string) bool {
					return len(key) == 6
				}), 300*time.Second).Return(errors.New("error")).Once()
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeEach()
			test.setup()
			key, err := s.notificationImpl.SendEmailRegisterCode(ctx, email)
			s.Equal(test.expectedKeyLen, len(key))
			s.Equal(test.expectedErr, err)

			mockCache.AssertExpectations(s.T())
			mockConfig.AssertExpectations(s.T())
			mockMail.AssertExpectations(s.T())
			mockLang.AssertExpectations(s.T())
		})
	}
}

func (s *AuthTestSuite) TestGetEmailRegisterCodeKey() {
	s.Equal(32, len(s.notificationImpl.getEmailRegisterCodeKey("hello@goravel.dev")))
}
