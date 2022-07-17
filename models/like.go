package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	Article   Article `gorm:"foreignKey:ArticleID"`
	ArticleID uint
	User      User `gorm:"foreignKey:UserID"`
	UserID    uint
}
