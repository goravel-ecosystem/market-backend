package services

import (
	"market.goravel.dev/package/app/models"
)

type Package interface {
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

func (r *PackageImpl) GetPackageByID(id string) (*models.Package, error) {
	return r.packageModel.GetPackageByID(id, []string{"id", "name", "user_id", "summary", "description", "link", "version", "last_updated_at"})
}
