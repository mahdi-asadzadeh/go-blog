package models

import (
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	FromUser   User `gorm:"foreignKey:FromUserID"`
	FromUserID uint
	ToUser     User `gorm:"foreignKey:ToUserID"`
	ToUserID   uint
}
