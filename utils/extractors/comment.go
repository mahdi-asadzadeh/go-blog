package extractors

import (
	"time"

	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetComment(co *models.Comment) map[string]interface{} {
	result := map[string]interface{}{
		"id":      co.ID,
		"content": co.Content,
		"user": map[string]interface{}{
			"id":         co.User.ID,
			"email":      co.User.Email,
			"first_name": co.User.FirstName,
			"last_name":  co.User.LastName,
			"bio":        co.User.Bio,
			"image":      co.User.Image,
		},
		"created_at": co.CreatedAt.UTC().Format(time.RFC1123),
		"updated_at": co.UpdatedAt.UTC().Format(time.RFC1123),
	}
	return result
}

func GetCommentList(cos []models.Comment) []interface{} {
	result := make([]interface{}, len(cos))
	for i := 0; i < len(cos); i++ {
		result[i] = GetComment(&cos[i])
	}
	return result
}
