package controllers

import (
	"context"
	"errors"
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/services"
	protopackage "market.goravel.dev/proto/package"
	utilsresponse "market.goravel.dev/utils/response"
)

type TagController struct {
	protopackage.UnimplementedTagServiceServer
	tagService services.Tag
}

func NewTagController() *TagController {
	return &TagController{
		tagService: services.NewTagImpl(),
	}
}

func (r *TagController) GetTags(ctx context.Context, req *protopackage.GetTagRequest) (*protopackage.GetTagResponse, error) {
	packageID := req.GetPackageId()
	userID := req.GetUserId()
	name := req.GetName()

	tags, err := r.tagService.GetTags(packageID, userID, name)
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return &protopackage.GetTagResponse{
			Status: utilsresponse.NewNotFoundStatus(errors.New(facades.Lang(ctx).Get("not_exist.user"))),
		}, nil
	}

	var tagsProto []*protopackage.Tag
	for _, tag := range tags {
		tagsProto = append(tagsProto, tag.ToProto())
	}

	return &protopackage.GetTagResponse{
		Status: utilsresponse.NewOkStatus(),
		Tags:   tagsProto,
	}, nil
}
