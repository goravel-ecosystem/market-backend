package services

import (
	"github.com/goravel/framework/facades"

	"market.goravel.dev/package/app/models"
)

type Tag interface {
	GetTags(packageID, userID, name string) ([]*models.Tag, error)
}

type TagImpl struct {
	tagModel models.TagInterface
}

func NewTagImpl() *TagImpl {
	return &TagImpl{
		tagModel: models.NewTag(),
	}
}

func (r *TagImpl) GetTags(packageID, userID, name string) ([]*models.Tag, error) {
	var tags []*models.Tag
	query := facades.Orm().Query()
	var total int64

	if packageID != "" {
		query = query.Join("JOIN package_tags ON package_tags.tag_id = tags.id").
			Where("package_tags.package_id = ?", packageID)
	}

	if userID != "" {
		query = query.Where("user_id", userID)
	}

	if name != "" {
		// fuzzy search
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Select([]string{"tags.id", "tags.name", "tags.user_id"}).Paginate(1, 10, &tags, &total); err != nil {
		return nil, err
	}

	return tags, nil
}
