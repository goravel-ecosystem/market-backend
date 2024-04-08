package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TagSuite struct {
	suite.Suite
	tag *Tag
}

func TestTagSuite(t *testing.T) {
	suite.Run(t, new(TagSuite))
}

func (s *TagSuite) SetupTest() {
	s.tag = NewTag()
}
