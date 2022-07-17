package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func FetchCategoryArticles(slug string) (models.Category, error) {
	database := infrastructure.GetDB()
	var category models.Category

	err := database.Model(&category).Where("slug = ?", slug).Preload("Articles").Preload("Articles.User").First(&category).Error
	return category, err
}

func FetchAllCategories() ([]models.Category, error) {
	database := infrastructure.GetDB()
	var categories []models.Category
	err := database.Find(&categories).Error
	return categories, err
}
