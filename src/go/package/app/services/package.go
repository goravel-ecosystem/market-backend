package services

import (
	"context"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	protouser "market.goravel.dev/proto/user"
	"market.goravel.dev/utils/errors"
)

type Package interface {
	GetPackages(query *protopackage.PackagesQuery, pagination *protobase.Pagination) ([]*models.Package, int64, error)
	GetPackageByID(id string) (*models.Package, error)
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

func (r *PackageImpl) GetPackages(query *protopackage.PackagesQuery, pagination *protobase.Pagination) (packages []*models.Package, total int64, err error) {
	const (
		categoryHot    = "hot"
		categoryNewest = "newest"
	)

	ormQuery := facades.Orm().Query()

	page := pagination.GetPage()
	limit := pagination.GetLimit()

	category := query.GetCategory()

	switch category {
	case categoryHot:
		ormQuery = ormQuery.OrderByDesc("view_count")
	case categoryNewest:
		ormQuery = ormQuery.OrderByDesc("created_at")
	}

	name := query.GetName()
	if name != "" {
		// fuzzy search
		ormQuery = ormQuery.Where("name LIKE ?", "%"+name+"%")
	}

	if err := ormQuery.With("Tags", func(query orm.Query) orm.Query {
		return query.Where("is_show = ?", "1").Select([]string{"id", "name"})
	}).Select([]string{"id", "name", "user_id", "summary", "link", "view_count"}).Paginate(int(page), int(limit), &packages, &total); err != nil {
		return nil, 0, errors.NewInternalServerError(err)
	}

	userIDs := make([]string, len(packages))
	for i, pkg := range packages {
		userIDs[i] = cast.ToString(pkg.UserID)
	}

	var users []*protouser.User
	if len(packages) > 0 {
		users, err = r.userService.GetUsers(context.Background(), userIDs)
		if err != nil {
			return nil, 0, errors.NewInternalServerError(err)
		}
	}

	userMap := make(map[string]*protouser.User)
	for _, user := range users {
		userMap[user.GetId()] = user
	}

	for _, pkg := range packages {
		pkg.User = userMap[cast.ToString(pkg.UserID)]
	}

	return packages, total, nil
}

func (r *PackageImpl) GetPackageByID(id string) (pkg *models.Package, err error) {
	pkg, err = r.packageModel.GetPackageByID(id, []string{"id", "name", "user_id", "summary", "description", "link", "version", "last_updated_at", "view_count"})
	if err != nil {
		return nil, err
	}

	if pkg.ID > 0 {
		user, err := r.userService.GetUser(context.Background(), pkg.UserID)
		pkg.User = user
		if err != nil {
			return nil, err
		}
	}

	return pkg, nil
}
