package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Email    string
	Password string
	Name     string
	Avatar   string
	Summary  string
	orm.SoftDeletes
}
