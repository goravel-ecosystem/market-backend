package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
	"github.com/spf13/cast"

	protopackage "market.goravel.dev/proto/package"
)

type TagInterface interface {
	//GetTags(packageID, userID, name string, fields []string) ([]*Tag, error)
}

type Tag struct {
	UUIDModel
	UserID      uint64
	Name        string
	Description string
	IsShow      uint
	orm.SoftDeletes
}

func NewTag() *Tag {
	return &Tag{}
}

func (r *Tag) ToProto() *protopackage.Tag {
	return &protopackage.Tag{
		Id:        cast.ToString(r.ID),
		UserId:    cast.ToString(r.UserID),
		Name:      r.Name,
		CreatedAt: r.CreatedAt.ToString(),
		UpdatedAt: r.UpdatedAt.ToString(),
		DeletedAt: carbon.FromStdTime(r.DeletedAt.Time).ToString(),
	}
}
