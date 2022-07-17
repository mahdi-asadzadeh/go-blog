package extractors

import (
	"net/http"
	"time"

	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/utils"
)

func GetArticleImage(im *models.ArticleImage) map[string]interface{} {
	result := map[string]interface{}{
		"path":       im.Path,
		"size":       im.Size,
		"created_at": im.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		"updated_at": im.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
	return result
}

func GetArticleImageList(ims []models.ArticleImage) []interface{} {
	result := make([]interface{}, len(ims))

	for i := 0; i < len(ims); i++ {
		result[i] = GetArticleImage(&ims[i])
	}
	return result
}

func GetArticleListPage(request *http.Request, articles []models.Article, page, page_size, count int) map[string]interface{} {
	resources := make([]interface{}, len(articles))
	for i := 0; i < len(articles); i++ {
		resources[i] = GetArticle(&articles[i])
	}
	return utils.CreatePagedResponse(request, resources, "articles", page, page_size, count)
}

func GetArticleList(articles []models.Article) []interface{} {
	result := make([]interface{}, len(articles))
	for i := 0; i < len(articles); i++ {
		result[i] = GetArticle(&articles[i])
	}
	return result
}

func GetArticle(data *models.Article) map[string]interface{} {
	result := map[string]interface{}{
		"id":          data.ID,
		"title":       data.Title,
		"slug":        data.Slug,
		"description": data.Description,
		"user": map[string]interface{}{
			"id":         data.User.ID,
			"email":      data.User.Email,
			"first_name": data.User.FirstName,
			"last_name":  data.User.LastName,
			"bio":        data.User.Bio,
			"image":      data.User.Email,
		},
		"created_at": data.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		"updated_at": data.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
	return result
}

func GetArticleDetail(data *models.Article) map[string]interface{} {
	result := map[string]interface{}{
		"id":          data.ID,
		"title":       data.Title,
		"slug":        data.Slug,
		"description": data.Description,
		"body":        data.Body,
		"user": map[string]interface{}{
			"id":         data.User.ID,
			"email":      data.User.Email,
			"first_name": data.User.FirstName,
			"last_name":  data.User.LastName,
			"bio":        data.User.Bio,
			"image":      data.User.Email,
		},
		"tags":       GetTagList(data.Tags),
		"categories": GetCategoryList(data.Categories),
		"created_at": data.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		"updated_at": data.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
	return result
}
