package controllers

import (
	"context"

	"github.com/goravel/framework/contracts/translation"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"

	protopackage "market.goravel.dev/proto/package"
	utilserrors "market.goravel.dev/utils/errors"
)

func validateCreatePackageRequest(ctx context.Context, req *protopackage.CreatePackageRequest) error {
	name := req.GetName()
	url := req.GetUrl()
	tags := req.GetTags()
	summery := req.GetSummary()
	description := req.GetDescription()
	userID := req.GetUserId()
	return validatePackageRequest(ctx, name, url, tags, summery, description, userID, req.GetLastUpdatedAt())
}

func validatePackageRequest(ctx context.Context, name, url string, tags []string, summary, description, userID, lastUpdatedAt string) error {
	translate := facades.Lang(ctx)
	if userID == "" {
		return utilserrors.NewBadRequest(translate.Get("required.user_id"))
	}
	if name == "" {
		return utilserrors.NewBadRequest(translate.Get("required.name"))
	}
	if len(name) > 100 {
		return utilserrors.NewBadRequest(translate.Get("max.name", translation.Option{
			Replace: map[string]string{
				"max": "100",
			},
		}))
	}

	if url == "" {
		return utilserrors.NewBadRequest(translate.Get("required.url"))
	}
	if len(url) > 100 {
		return utilserrors.NewBadRequest(translate.Get("max.url", translation.Option{
			Replace: map[string]string{
				"max": "100",
			},
		}))
	}

	if len(tags) > 10 {
		return utilserrors.NewBadRequest(translate.Get("max.tags", translation.Option{
			Replace: map[string]string{
				"max": "10",
			},
		}))
	}

	if len(summary) > 200 {
		return utilserrors.NewBadRequest(translate.Get("max.summary", translation.Option{
			Replace: map[string]string{
				"max": "200",
			},
		}))
	}

	if len(description) > 10000 {
		return utilserrors.NewBadRequest(translate.Get("max.description", translation.Option{
			Replace: map[string]string{
				"max": "10000",
			},
		}))
	}

	if lastUpdatedAt == "" {
		return utilserrors.NewBadRequest(translate.Get("required.last_updated_at"))
	}
	if carbon.Parse(lastUpdatedAt).IsZero() {
		return utilserrors.NewBadRequest(translate.Get("invalid.last_updated_at"))
	}

	return nil
}

func validateUpdatePackageRequest(ctx context.Context, req *protopackage.UpdatePackageRequest) error {
	name := req.GetName()
	url := req.GetUrl()
	tags := req.GetTags()
	summery := req.GetSummary()
	description := req.GetDescription()
	userID := req.GetUserId()
	id := req.GetId()
	if id == "" {
		return utilserrors.NewBadRequest(facades.Lang(ctx).Get("required.id"))
	}

	return validatePackageRequest(ctx, name, url, tags, summery, description, userID, req.GetLastUpdatedAt())
}
