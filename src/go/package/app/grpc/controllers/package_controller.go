package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
	"market.goravel.dev/utils/response"
)

type PackageController struct {
	protopackage.UnimplementedPackageServiceServer
	packageService services.Package
}

func NewPackageController() *PackageController {
	return &PackageController{
		packageService: services.NewPackageImpl(),
	}
}

func (r *PackageController) GetPackage(ctx context.Context, req *protopackage.GetPackageRequest) (*protopackage.GetPackageResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		return &protopackage.GetPackageResponse{
			Status: response.NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.user_id"))),
		}, nil
	}
	packageID := req.GetId()
	if packageID == "" {
		return &protopackage.GetPackageResponse{
			Status: response.NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.package_id"))),
		}, nil
	}

	packageData, err := r.packageService.GetPackageByID(packageID)
	if err != nil {
		if errors.Is(err, orm.ErrRecordNotFound) {
			return &protopackage.GetPackageResponse{
				Status: response.NewNotFoundStatus(errors.New(facades.Lang(ctx).Get("not_exist.package"))),
			}, nil
		}
		return nil, err
	}

	return &protopackage.GetPackageResponse{
		Status:  response.NewOkStatus(),
		Package: packageData.ToProto(),
	}, nil
}
