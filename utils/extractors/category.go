package extractors

import (
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetCategory(cat models.Category) map[string]interface{} {
	result := map[string]interface{}{
		"id":          cat.ID,
		"name":        cat.Name,
		"slug":        cat.Slug,
		"description": cat.Description,
	}
	return result
}

func GetCategoryList(cats []models.Category) []interface{} {
	result := make([]interface{}, len(cats))
	for i := 0; i < len(cats); i++ {
		result[i] = GetCategory(cats[i])
	}
	return result
}
