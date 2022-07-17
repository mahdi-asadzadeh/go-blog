package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"unique_index"`
	Slug        string `gorm:"unique_index"`
	Description string
	Articles    []Article `gorm:"many2many:articles_categories;"`
	IsNewRecord bool      `gorm:"-;default:false"`
}

func (a *Category) BeforeSave() (err error) {
	a.Slug = slug.Make(a.Name)
	return
}
