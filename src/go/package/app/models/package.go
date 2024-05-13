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
	AttachTags(pkg *Package, tags []string) error
	GetPackageByID(id string, fields []string) (*Package, error)
	UpdatePackage(pkg *Package) error
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

func (r *Package) AttachTags(pkg *Package, tags []string) error {
	tagsAny := make([]any, len(tags))
	for i, tag := range tags {
		tagsAny[i] = tag
	}
	existTags := make([]*Tag, 0, len(tags))
	if err := facades.Orm().Query().WhereIn("name", tagsAny).Find(&existTags); err != nil {
		return errors.NewInternalServerError(err)
	}
	existTagMap := make(map[string]bool, len(existTags))
	for _, tag := range existTags {
		existTagMap[tag.Name] = true
	}

	newTags := make([]*Tag, 0, len(tags))
	for _, tag := range tags {
		if !existTagMap[tag] {
			newTag := Tag{
				Name:   tag,
				UserID: pkg.UserID,
				IsShow: 1,
			}
			newTag.ID = newTag.GetID()
			newTags = append(newTags, &newTag)
		}
	}

	if len(newTags) > 0 {
		if err := facades.Orm().Query().Create(newTags); err != nil {
			return errors.NewInternalServerError(err)
		}
		existTags = append(existTags, newTags...)
	}

	if err := facades.Orm().Query().Model(pkg).Association("Tags").Replace(existTags); err != nil {
		return errors.NewInternalServerError(err)
	}

	return nil
}

func (r *Package) GetPackageByID(id string, fields []string) (*Package, error) {
	var packageModel Package

	query := facades.Orm().Query().Where("id", id).With("Tags", func(query contractsorm.Query) contractsorm.Query {
		return query.Where("is_show = ?", "1").Select([]string{"id", "name"})
	})

	if len(fields) > 0 {
		query = query.Select(fields)
	}

	if err := query.First(&packageModel); err != nil {
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

func (r *Package) UpdatePackage(pkg *Package) error {
	if err := facades.Orm().Query().Save(pkg); err != nil {
		return errors.NewInternalServerError(err)
	}

	return nil
}
