package services

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/models"
	protobase "market.goravel.dev/proto/base"
)

type Tag interface {
	GetTags(packageID, name string, pagination *protobase.Pagination) ([]*models.Tag, int64, error)
}

type TagImpl struct {
}

func NewTagImpl() *TagImpl {
	return &TagImpl{}
}

func (r *TagImpl) GetTags(packageID, name string, pagination *protobase.Pagination) ([]*models.Tag, int64, error) {
	var tags []*models.Tag
	query := facades.Orm().Query()
	var total int64

	page := pagination.GetPage()
	limit := pagination.GetLimit()

	if packageID != "" {
		var tagIDs []any
		if err := facades.Orm().Query().Table("package_tags").Where("package_id = ?", packageID).Pluck("tag_id", &tagIDs); err != nil {
			return nil, 0, err
		}

		query = query.WhereIn("id", tagIDs)
	}

	if name != "" {
		// fuzzy search
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if err := query.Select([]string{"id", "name"}).Where("is_show = ?", "1").Paginate(int(page), int(limit), &tags, &total); err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}
