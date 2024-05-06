package controllers

import (
	"context"

	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
	protouser "market.goravel.dev/proto/user"
	utilserrors "market.goravel.dev/utils/errors"
	utilspagination "market.goravel.dev/utils/pagination"
	utilsresponse "market.goravel.dev/utils/response"
)

type PackageController struct {
	protopackage.UnimplementedPackageServiceServer
	packageService services.Package
	tagService     services.Tag
	userService    services.User
}

func NewPackageController() *PackageController {
	return &PackageController{
		packageService: services.NewPackageImpl(),
		tagService:     services.NewTagImpl(),
		userService:    services.NewUserImpl(),
	}
}

func (r *PackageController) GetPackage(ctx context.Context, req *protopackage.GetPackageRequest) (*protopackage.GetPackageResponse, error) {
	packageID := req.GetId()
	if packageID == "" {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.package_id"))
	}

	pkg, err := r.packageService.GetPackageByID(packageID)
	if err != nil {
		return nil, err
	}

	if pkg.ID == 0 {
		return nil, utilserrors.NewNotFound(facades.Lang(ctx).Get("not_exist.package"))
	}

	user, err := r.userService.GetUser(ctx, pkg.UserID)
	if err != nil {
		return nil, err
	}

	tags, err := r.GetTags(ctx, &protopackage.GetTagsRequest{
		Query: &protopackage.TagsQuery{
			PackageId: packageID,
		},
	})
	if err != nil {
		return nil, err
	}

	pkgProto := pkg.ToProto()
	pkgProto.User = user
	pkgProto.Tags = tags.GetTags()

	return &protopackage.GetPackageResponse{
		Status:  utilsresponse.NewOkStatus(),
		Package: pkgProto,
	}, nil
}

func (r *PackageController) GetPackages(ctx context.Context, req *protopackage.GetPackagesRequest) (*protopackage.GetPackagesResponse, error) {
	query := req.GetQuery()
	pagination := req.GetPagination()

	if pagination == nil {
		pagination = utilspagination.Default()
	}

	if pagination.GetPage() <= 0 {
		pagination.Page = 1
	}

	if pagination.GetLimit() <= 0 {
		pagination.Limit = 10
	}

	packages, total, err := r.packageService.GetPackages(query, pagination)
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, len(packages))
	for i, pkg := range packages {
		userIDs[i] = cast.ToString(pkg.UserID)
	}

	var users []*protouser.User
	if len(packages) > 0 {
		users, err = r.userService.GetUsers(ctx, userIDs)
		if err != nil {
			return nil, err
		}
	}

	userMap := make(map[string]*protouser.User)
	for _, user := range users {
		userMap[user.GetId()] = user
	}

	packagesProto := make([]*protopackage.Package, 0, len(packages))
	for _, pkg := range packages {
		pkgProto := pkg.ToProto()
		if user, ok := userMap[cast.ToString(pkg.UserID)]; ok {
			pkgProto.User = user
		}
		packagesProto = append(packagesProto, pkgProto)
	}

	return &protopackage.GetPackagesResponse{
		Status:   utilsresponse.NewOkStatus(),
		Packages: packagesProto,
		Total:    total,
	}, nil
}

func (r *PackageController) GetTags(_ context.Context, req *protopackage.GetTagsRequest) (*protopackage.GetTagsResponse, error) {
	query := req.GetQuery()
	packageID := query.GetPackageId()
	name := query.GetName()
	pagination := req.GetPagination()

	if pagination == nil {
		pagination = utilspagination.Default()
	}

	if pagination.GetPage() <= 0 {
		pagination.Page = 1
	}

	if pagination.GetLimit() <= 0 {
		pagination.Limit = 10
	}

	tags, total, err := r.tagService.GetTags(packageID, name, pagination)
	if err != nil {
		return nil, err
	}

	tagsProto := make([]*protopackage.Tag, 0)
	for _, tag := range tags {
		tagsProto = append(tagsProto, tag.ToProto())
	}

	return &protopackage.GetTagsResponse{
		Status: utilsresponse.NewOkStatus(),
		Tags:   tagsProto,
		Total:  total,
	}, nil
}
