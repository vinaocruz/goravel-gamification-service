package models

import (
	"github.com/goravel/framework/database/orm"
)

type PlayerEvents struct {
	orm.Model
	PlayerID uint `gorm:"primaryKey"`
	EventID  uint `gorm:"primaryKey"`
	Points   int
}
