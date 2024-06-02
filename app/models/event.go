package models

import (
	"github.com/goravel/framework/database/orm"
)

type Event struct {
	orm.Model
	Name   string `form:"name"`
	Points int    `form:"point"`
	orm.SoftDeletes
}
