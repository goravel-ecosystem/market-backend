package controllers

import (
	"context"

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
}

func NewPackageController() *PackageController {
	return &PackageController{
		packageService: services.NewPackageImpl(),
		tagService:     services.NewTagImpl(),
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

	return &protopackage.GetPackageResponse{
		Status:  utilsresponse.NewOkStatus(),
		Package: pkg.ToProto(),
	}, nil
}

func (r *PackageController) GetPackages(_ context.Context, req *protopackage.GetPackagesRequest) (*protopackage.GetPackagesResponse, error) {
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

	packagesProto := make([]*protopackage.Package, 0, len(packages))
	for _, pkg := range packages {
		packagesProto = append(packagesProto, pkg.ToProto())
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
