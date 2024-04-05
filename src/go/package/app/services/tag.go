package services

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
)

type Tag interface {
	GetTags(packageID, name string, pagination *protobase.Pagination, total *int64) ([]*models.Tag, error)
}

type TagImpl struct {
	tagModel models.TagInterface
}

func NewTagImpl() *TagImpl {
	return &TagImpl{
		tagModel: models.NewTag(),
	}
}

func (r *TagImpl) GetTags(packageID, name string, pagination *protobase.Pagination, total *int64) ([]*models.Tag, error) {
	var tags []*models.Tag
	query := facades.Orm().Query()

	if packageID != "" {
		query = query.Join("JOIN package_tags ON package_tags.tag_id = tags.id").
			Where("package_tags.package_id = ?", packageID)
	}

	if name != "" {
		// fuzzy search
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	page := pagination.GetPage()
	limit := pagination.GetLimit()
	if limit <= 0 {
		limit = 20
	}

	if page <= 0 {
		page = 1
	}
	if err := query.Select([]string{"tags.id", "tags.name", "tags.user_id"}).Paginate(int(page), int(limit), &tags, total); err != nil {
		return nil, err
	}

	return tags, nil
}
