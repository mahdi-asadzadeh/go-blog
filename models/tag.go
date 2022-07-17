package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name        string `gorm:"unique_index"`
	Slug        string `gorm:"unique_index"`
	Description string
	Articles    []Article `gorm:"many2many:articles_tags"`
}

func (a *Tag) BeforeSave() (err error) {
	a.Slug = slug.Make(a.Name)
	return
}

type ArticleTag struct {
	Tag       Tag `gorm:"foreignkey:TagID"`
	TagID     uint
	Article   Article `gorm:"foreignkey:ArticleID"`
	ArticleID uint
}

func (ArticleTag) TableName() string {
	return "articles_tags"
}
