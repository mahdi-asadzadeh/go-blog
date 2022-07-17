package extractors

import (
	"time"

	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetLike(l *models.Like) map[string]interface{} {
	result := map[string]interface{}{
		"user": map[string]interface{}{
			"id":         l.User.ID,
			"email":      l.User.Email,
			"first_name": l.User.FirstName,
			"last_name":  l.User.LastName,
			"bio":        l.User.Bio,
			"image":      l.User.Email,
		},
		"created_at": l.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		"updated_at": l.UpdatedAt.UTC().Format(time.RFC3339Nano),
	}
	return result
}

func GetLikeList(ls []models.Like) []interface{} {
	result := make([]interface{}, len(ls))
	for i := 0; i < len(ls); i++ {
		result[i] = GetLike(&ls[i])
	}
	return result
}
