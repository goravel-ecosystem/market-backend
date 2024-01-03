package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type Package struct {
	orm.Model
	UserID        uint
	Name          string
	Summary       string
	Description   string
	Link          string
	Version       string
	LastUpdatedAt carbon.DateTime
	orm.SoftDeletes
}
