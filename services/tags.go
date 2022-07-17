package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func FetchAllTags() ([]models.Tag, error) {
	database := infrastructure.GetDB()

	var tags []models.Tag
	err := database.Model(&tags).Find(&tags).Error
	return tags, err
}
