package services

import (
	"testing"

	"github.com/stretchr/testify/suite"

	mocksmodels "market.goravel.dev/package/app/mocks/models"
)

type UserTestSuite struct {
	suite.Suite
	mockTag *mocksmodels.TagInterface
	tagImpl *TagImpl
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupTest() {
	s.mockTag = new(mocksmodels.TagInterface)
	s.tagImpl = &TagImpl{
		tagModel: s.mockTag,
	}
}

func (s *UserTestSuite) TestGetTags() {
	// TODO: Implement me
}
