package models

import (
	"github.com/goravel/framework/database/orm"
)

type Tag struct {
	orm.Model
	Name string
	orm.SoftDeletes
}
