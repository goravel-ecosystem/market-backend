package controllers

import (
	"context"
	"errors"

	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
	utilsresponse "market.goravel.dev/utils/response"
)

type PackageController struct {
	protopackage.UnimplementedPackageServiceServer
	tagService services.Tag
}

func NewPackageController() *PackageController {
	return &PackageController{
		tagService: services.NewTagImpl(),
	}
}

func (r *PackageController) GetTags(ctx context.Context, req *protopackage.GetTagsRequest) (*protopackage.GetTagsResponse, error) {
	query := req.GetQuery()
	packageID := query.GetPackageId()
	name := query.GetName()
	pagination := req.GetPagination()
	var total int64

	tags, err := r.tagService.GetTags(packageID, name, pagination, &total)
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return &protopackage.GetTagsResponse{
			Status: utilsresponse.NewNotFoundStatus(errors.New(facades.Lang(ctx).Get("not_exist.tags"))),
		}, nil
	}

	var tagsProto []*protopackage.Tag
	for _, tag := range tags {
		tagsProto = append(tagsProto, tag.ToProto())
	}

	return &protopackage.GetTagsResponse{
		Status: utilsresponse.NewOkStatus(),
		Tags:   tagsProto,
		Total:  total,
	}, nil
}
