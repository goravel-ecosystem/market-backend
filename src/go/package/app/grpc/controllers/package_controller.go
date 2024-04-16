package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
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
	userID := req.GetUserId()
	if userID == "" {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.user_id"))
	}
	packageID := req.GetId()
	if packageID == "" {
		return nil, utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.package_id"))
	}

	pkg, err := r.packageService.GetPackageByID(packageID)
	if err != nil {
		if errors.Is(err, orm.ErrRecordNotFound) {
			return nil, utilserrors.NewNotFound(facades.Lang(ctx).Get("not_exist.package"))
		}
		return nil, err
	}

	user, err := r.userService.GetUser(ctx, pkg.UserID)
	if err != nil {
		return nil, err
	}

	pkgProto := pkg.ToProto()
	pkgProto.User = user

	tags, err := r.GetTags(ctx, &protopackage.GetTagsRequest{
		Query: &protopackage.TagsQuery{
			PackageId: packageID,
		},
	})

	if err != nil {
		return nil, err
	}

	pkgProto.Tags = tags.GetTags()

	return &protopackage.GetPackageResponse{
		Status:  utilsresponse.NewOkStatus(),
		Package: pkgProto,
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
