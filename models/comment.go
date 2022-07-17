package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string  `gorm:"size:2048"`
	Article   Article `gorm:"foreignKey:ArticleID"`
	ArticleID uint
	User      User `gorm:"foreignKey:UserID"`
	UserID    uint
}
