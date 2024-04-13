package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
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
	userID := req.GetUserId()
	if userID == "" {
		return &protopackage.GetPackageResponse{
			Status: utilsresponse.NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.user_id"))),
		}, nil
	}
	packageID := req.GetId()
	if packageID == "" {
		return &protopackage.GetPackageResponse{
			Status: utilsresponse.NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.package_id"))),
		}, nil
	}

	packageData, err := r.packageService.GetPackageByID(packageID)
	if err != nil {
		if errors.Is(err, orm.ErrRecordNotFound) {
			return &protopackage.GetPackageResponse{
				Status: utilsresponse.NewNotFoundStatus(errors.New(facades.Lang(ctx).Get("not_exist.package"))),
			}, nil
		}
		return nil, err
	}

	return &protopackage.GetPackageResponse{
		Status:  utilsresponse.NewOkStatus(),
		Package: packageData.ToProto(),
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
