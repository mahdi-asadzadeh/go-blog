package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func FetchCommentByID(id int) (models.Comment, error) {
	database := infrastructure.GetDB()
	var comment models.Comment
	err := database.Model(&comment).Where("id = ?", id).Preload("User").First(&comment).Error
	return comment, err
}
