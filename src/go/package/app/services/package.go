package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	"market.goravel.dev/utils/errors"
)

type Package interface {
	GetPackages(query *protopackage.PackagesQuery, pagination *protobase.Pagination) ([]*models.Package, int64, error)
	GetPackageByID(id string) (*models.Package, error)
}

type PackageImpl struct {
	packageModel models.PackageInterface
}

func NewPackageImpl() *PackageImpl {
	return &PackageImpl{
		packageModel: models.NewPackage(),
	}
}

func (r *PackageImpl) GetPackages(query *protopackage.PackagesQuery, pagination *protobase.Pagination) ([]*models.Package, int64, error) {
	var packages []*models.Package
	ormQuery := facades.Orm().Query()
	var total int64

	page := pagination.GetPage()
	limit := pagination.GetLimit()

	name := query.GetName()
	if name != "" {
		// fuzzy search
		ormQuery = ormQuery.Where("name LIKE ?", "%"+name+"%")
	}

	if err := ormQuery.With("Tags", func(query orm.Query) orm.Query {
		return query.Where("is_show = ?", "1").Select([]string{"id", "name"})
	}).Select([]string{"id", "name", "user_id", "summary", "link"}).Paginate(int(page), int(limit), &packages, &total); err != nil {
		return nil, 0, errors.NewInternalServerError(err)
	}

	return packages, total, nil
}

func (r *PackageImpl) GetPackageByID(id string) (*models.Package, error) {
	return r.packageModel.GetPackageByID(id, []string{"id", "name", "user_id", "summary", "description", "link", "version", "last_updated_at"})
}
