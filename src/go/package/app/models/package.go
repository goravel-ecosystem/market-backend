package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/spf13/cast"

	protopackage "market.goravel.dev/proto/package"
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
	Version       string
	LastUpdatedAt carbon.DateTime
	orm.SoftDeletes
}

func NewPackage() *Package {
	return &Package{}
}

func (r *Package) GetPackageByID(id string, fields []string) (*Package, error) {
	var packageModel Package
	if err := facades.Orm().Query().Where("id", id).Select(fields).FirstOrFail(&packageModel); err != nil {
		return nil, err
	}

	return &packageModel, nil
}

func (r *Package) ToProto() *protopackage.Package {
	return &protopackage.Package{
		Id:            cast.ToString(r.ID),
		UserId:        cast.ToString(r.UserID),
		Name:          r.Name,
		Summary:       r.Summary,
		Description:   r.Description,
		Link:          r.Link,
		Version:       r.Version,
		LastUpdatedAt: r.LastUpdatedAt.ToString(),
		CreatedAt:     r.CreatedAt.ToString(),
		UpdatedAt:     r.UpdatedAt.ToString(),
		DeletedAt:     carbon.FromStdTime(r.DeletedAt.Time).ToString(),
	}
}
