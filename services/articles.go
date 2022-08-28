package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func DeleteArticleIfOwnerOrAdmin(user *models.User, slug string) error {
	database := infrastructure.GetDB()
	var article models.Article
	err := database.Model(&article).Where("slug = ?", slug).Select("user_id").First(&article).Error
	if err != nil {
		return err
	}
	if user.ID == article.UserID {
		article.Slug = slug
		err = database.Where("slug = ?", slug).Delete(&article).Error
	}
	return err
}

func FetchArticleId(slug string) (uint, error) {
	articleId := -1
	database := infrastructure.GetDB()
	err := database.Model(&models.Article{}).Where("slug = ?", slug).Select("id").Row().Scan(&articleId)
	return uint(articleId), err
}

func FetchArticleImages(articleID uint) ([]models.ArticleImage, error) {
	database := infrastructure.GetDB()
	var images []models.ArticleImage

	err := database.Model(&images).Where("article_id = ?", articleID).Find(&images).Error
	return images, err
}

func FetchArticleDetails(condition interface{}) (models.Article, error) {
	database := infrastructure.GetDB()

	var article models.Article
	err := database.Where(condition).Preload("Tags").Preload("User").Preload("Categories").First(&article).Error
	return article, err
}

func FetchArticlesPage(page int, pageSize int) ([]models.Article, int, error) {
	db := infrastructure.GetDB()

	var count int64
	var articles []models.Article
	offSet := (page - 1) * pageSize

	db.Model(&articles).Count(&count)
	err := db.Offset(offSet).Limit(pageSize).Preload("User").Find(&articles).Error

	return articles, int(count), err
}

func CreateOneArticle(article *models.Article) error {
	database := infrastructure.GetDB()
	err := article.BeforeSave()
	if err != nil {
		return err
	}
	err = database.Create(article).Error
	return err
}

func CreateOneImageArticle(img *models.ArticleImage) error {
	database := infrastructure.GetDB()
	err := database.Create(img).Error
	return err
}
