package services

import (
	"context"

	"market.goravel.dev/package/app/models"
)

type Package interface {
	GetPackageByID(ctx context.Context, id string) (*models.Package, error)
}

type PackageImpl struct {
	packageModel models.PackageInterface
	userService  User
}

func NewPackageImpl() *PackageImpl {
	return &PackageImpl{
		packageModel: models.NewPackage(),
		userService:  NewUserImpl(),
	}
}

func (r *PackageImpl) GetPackageByID(ctx context.Context, id string) (*models.Package, error) {
	pkg, err := r.packageModel.GetPackageByID(id, []string{"id", "name", "user_id", "summary", "description", "link", "version", "last_updated_at"})
	if err != nil {
		return nil, err
	}
	user, err := r.userService.GetUser(ctx, pkg.UserID)
	//TODO Add user to pkg

	return pkg, err
}
