package models

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/spf13/cast"

	protopackage "market.goravel.dev/proto/package"
	protouser "market.goravel.dev/proto/user"
	"market.goravel.dev/utils/errors"
)

type PackageInterface interface {
	GetPackageByID(id string, fields []string) (*Package, error)
}

type Package struct {
	UUIDModel
	UserID        uint64
	Name          string
	Summary       string
	Description   string
	Link          string
	Cover         string
	Version       string
	ViewCount     uint32
	IsPublic      int32
	IsApproved    int32
	LastUpdatedAt carbon.DateTime
	Tags          []*Tag          `gorm:"many2many:package_tags;"`
	User          *protouser.User `gorm:"-"`
	orm.SoftDeletes
}

func NewPackage() *Package {
	return &Package{}
}

func (r *Package) GetPackageByID(id string, fields []string) (*Package, error) {
	var packageModel Package
	if err := facades.Orm().Query().Where("id", id).With("Tags", func(query contractsorm.Query) contractsorm.Query {
		return query.Where("is_show = ?", "1").Select([]string{"id", "name"})
	}).Select(fields).First(&packageModel); err != nil {
		return nil, errors.NewInternalServerError(err)
	}

	return &packageModel, nil
}

func (r *Package) ToProto() *protopackage.Package {
	tagsProto := make([]*protopackage.Tag, 0)
	for _, tag := range r.Tags {
		tagsProto = append(tagsProto, tag.ToProto())
	}

	var isPublic bool
	if r.IsPublic == 2 {
		isPublic = true
	}

	return &protopackage.Package{
		Id:            cast.ToString(r.ID),
		UserId:        cast.ToString(r.UserID),
		Name:          r.Name,
		Summary:       r.Summary,
		Description:   r.Description,
		Link:          r.Link,
		Version:       r.Version,
		ViewCount:     r.ViewCount,
		LastUpdatedAt: r.LastUpdatedAt.ToString(),
		CreatedAt:     r.CreatedAt.ToString(),
		UpdatedAt:     r.UpdatedAt.ToString(),
		Tags:          tagsProto,
		User:          r.User,
		IsPublic:      isPublic,
		Cover:         r.Cover,
	}
}
