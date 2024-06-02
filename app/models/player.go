package models

import (
	"github.com/goravel/framework/database/orm"
)

type Player struct {
	orm.Model
	Name   string   `form:"name"`
	Email  string   `form:"email"`
	Events []*Event `gorm:"many2many:player_events"`
}
