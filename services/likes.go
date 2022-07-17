package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func IsLikeBy(article *models.Article, user *models.User) bool {
	database := infrastructure.GetDB()
	var favorite models.Like
	database.Where(models.Like{
		ArticleID: article.ID,
		UserID:    user.ID,
	}).First(&favorite)
	return favorite.ID != 0
}

func CreateOneLike(like *models.Like) error {
	database := infrastructure.GetDB()
	err := database.Create(like).Error
	return err
}
