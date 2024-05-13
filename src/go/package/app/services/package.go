package services

import (
	"context"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/spf13/cast"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
	protopackage "market.goravel.dev/proto/package"
	protouser "market.goravel.dev/proto/user"
	"market.goravel.dev/utils/errors"
)

type Package interface {
	CreatePackage(req *protopackage.CreatePackageRequest) (*models.Package, error)
	GetPackages(query *protopackage.PackagesQuery, pagination *protobase.Pagination) ([]*models.Package, int64, error)
	GetPackageByID(id string) (*models.Package, error)
	UpdatePackage(ctx context.Context, req *protopackage.UpdatePackageRequest) (*models.Package, error)
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

func (r *PackageImpl) CreatePackage(req *protopackage.CreatePackageRequest) (*models.Package, error) {
	pkg := models.Package{
		UserID:        cast.ToUint64(req.GetUserId()),
		Name:          req.GetName(),
		Summary:       req.GetSummary(),
		Description:   req.GetDescription(),
		Link:          req.GetUrl(),
		Cover:         req.GetCover(),
		Version:       req.GetVersion(),
		IsPublic:      req.GetIsPublic(),
		LastUpdatedAt: carbon.DateTime{Carbon: carbon.Parse(req.GetLastUpdatedAt())},
	}

	pkg.ID = pkg.GetID()

	if err := facades.Orm().Query().Create(&pkg); err != nil {
		return nil, errors.NewInternalServerError(err)
	}

	// Tags
	tags := req.GetTags()
	if len(tags) > 0 {
		tagsAny := make([]any, len(tags))
		for i, tag := range tags {
			tagsAny[i] = tag
		}
		existTags := make([]*models.Tag, 0, len(tags))
		if err := facades.Orm().Query().WhereIn("name", tagsAny).Find(&existTags); err != nil {
			return nil, errors.NewInternalServerError(err)
		}
		existTagMap := make(map[string]bool, len(existTags))
		for _, tag := range existTags {
			existTagMap[tag.Name] = true
		}

		newTags := make([]*models.Tag, 0, len(tags))
		for _, tag := range tags {
			if !existTagMap[tag] {
				newTag := models.Tag{
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
				return nil, errors.NewInternalServerError(err)
			}
			existTags = append(existTags, newTags...)
		}

		if err := facades.Orm().Query().Model(&pkg).Association("Tags").Replace(existTags); err != nil {
			return nil, errors.NewInternalServerError(err)
		}
	}

	return &pkg, nil
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

	userID := query.GetUserId()
	if userID != "" {
		ormQuery = ormQuery.Where("user_id = ?", userID)
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

func (r *PackageImpl) UpdatePackage(ctx context.Context, req *protopackage.UpdatePackageRequest) (*models.Package, error) {
	pkg, err := r.packageModel.GetPackageByID(req.GetId(), []string{})
	if err != nil {
		return nil, err
	}

	if pkg.ID == 0 {
		return nil, errors.NewNotFound(facades.Lang(ctx).Get("not_exist.package"))
	}

	if pkg.UserID != cast.ToUint64(req.GetUserId()) {
		return nil, errors.NewUnauthorized(facades.Lang(ctx).Get("forbidden.update_package"))
	}

	pkg.Name = req.GetName()
	pkg.Summary = req.GetSummary()
	pkg.Description = req.GetDescription()
	pkg.Link = req.GetUrl()
	pkg.Cover = req.GetCover()
	pkg.Version = req.GetVersion()
	pkg.IsPublic = req.GetIsPublic()
	pkg.LastUpdatedAt = carbon.DateTime{Carbon: carbon.Parse(req.GetLastUpdatedAt())}

	if err := r.packageModel.UpdatePackage(pkg); err != nil {
		return nil, err
	}

	// Tags
	tags := req.GetTags()
	if len(tags) > 0 {
		if err := r.packageModel.AttachTags(pkg, tags); err != nil {
			return nil, err
		}
	}

	return pkg, nil
}
