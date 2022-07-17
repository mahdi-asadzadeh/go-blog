package extractors

import (
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetTag(t models.Tag) map[string]interface{} {
	result := map[string]interface{}{
		"id":          t.ID,
		"name":        t.Name,
		"slug":        t.Slug,
		"description": t.Description,
	}
	return result
}

func GetTagList(ts []models.Tag) []interface{} {
	result := make([]interface{}, len(ts))
	for i := 0; i < len(ts); i++ {
		result[i] = GetTag(ts[i])
	}
	return result
}
