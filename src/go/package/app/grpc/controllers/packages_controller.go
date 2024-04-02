package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
)

type PackagesController struct {
	protopackage.UnimplementedPackageServiceServer
	packageService services.Package
}

func NewPackagesController() *PackagesController {
	return &PackagesController{
		packageService: services.NewPackageImpl(),
	}
}

func (r *PackagesController) GetPackage(ctx context.Context, req *protopackage.GetPackageRequest) (*protopackage.GetPackageResponse, error) {
	userID := req.GetUserId()
	if userID == "" {
		return &protopackage.GetPackageResponse{
			Status: NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.user_id"))),
		}, nil
	}
	packageId := req.GetId()
	if packageId == "" {
		return &protopackage.GetPackageResponse{
			Status: NewBadRequestStatus(errors.New(facades.Lang(ctx).Get("required.package_id"))),
		}, nil
	}

	packageData, err := r.packageService.GetPackageByID(packageId)
	if err != nil {
		return nil, err
	}

	return &protopackage.GetPackageResponse{
		Status:  NewOkStatus(),
		Package: packageData.ToProto(),
	}, nil
}
