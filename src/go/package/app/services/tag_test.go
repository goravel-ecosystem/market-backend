package services

import (
	"testing"

	"github.com/stretchr/testify/suite"

	mocksmodels "market.goravel.dev/package/app/mocks/models"
)

type TagTestSuite struct {
	suite.Suite
	mockTag *mocksmodels.TagInterface
	tagImpl *TagImpl
}

func TestTagTestSuite(t *testing.T) {
	suite.Run(t, new(TagTestSuite))
}

func (s *TagTestSuite) SetupTest() {
	s.mockTag = new(mocksmodels.TagInterface)
	s.tagImpl = &TagImpl{
		tagModel: s.mockTag,
	}
}

func (s *TagTestSuite) TestGetTags() {
	// TODO: Implement me
}
