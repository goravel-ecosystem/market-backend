package models

import (
	"testing"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
	"github.com/stretchr/testify/suite"

	protopackage "market.goravel.dev/proto/package"
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

func (s *TagSuite) TestToProto() {
	var (
		id        = 1
		userID    = 1
		name      = "Go"
		createAt  = carbon.DateTime{}
		updatedAt = carbon.DateTime{}
	)

	tag := Tag{
		UUIDModel: UUIDModel{
			ID: uint64(id),
			Timestamps: orm.Timestamps{
				CreatedAt: createAt,
				UpdatedAt: updatedAt,
			},
		},
		UserID: uint64(userID),
		Name:   name,
	}

	s.Equal(&protopackage.Tag{
		Id:        "1",
		UserId:    "1",
		Name:      name,
		CreatedAt: createAt.ToString(),
		UpdatedAt: updatedAt.ToString(),
	}, tag.ToProto())
}
