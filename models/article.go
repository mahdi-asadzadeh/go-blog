package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ArticleImage struct {
	gorm.Model
	Path      string `gorm:"uniqueIndex"`
	Size      int64
	Article   Article `gorm:"foreignKey:ArticleID"`
	ArticleID uint
}

type Article struct {
	gorm.Model
	Slug        string `gorm:"uniqueIndex"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	UserID      uint
	User        User
	Categories  []Category `gorm:"many2many:articles_categories"`
	Tags        []Tag      `gorm:"many2many:articles_tags"`
}

func (a *Article) BeforeSave() (err error) {
	a.Slug = slug.Make(a.Title)
	return nil
}
