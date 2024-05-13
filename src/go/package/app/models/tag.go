package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/spf13/cast"

	protopackage "market.goravel.dev/proto/package"
)

type Tag struct {
	UUIDModel
	UserID uint64
	Name   string
	IsShow uint
	orm.SoftDeletes
}

func NewTag() *Tag {
	return &Tag{}
}

func (r *Tag) ToProto() *protopackage.Tag {
	var userID string
	if r.UserID != 0 {
		userID = cast.ToString(r.UserID)
	}

	return &protopackage.Tag{
		Id:     cast.ToString(r.ID),
		UserId: userID,
		Name:   r.Name,
	}
}
