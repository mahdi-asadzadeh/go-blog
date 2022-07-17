package extractors

import (
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetUserFollower(f *models.Relation) map[string]interface{} {
	result := map[string]interface{}{
		"id":         f.ToUser.ID,
		"email":      f.ToUser.Email,
		"first_name": f.ToUser.FirstName,
		"last_name":  f.ToUser.LastName,
		"bio":        f.ToUser.Bio,
		"image":      f.ToUser.Email,
	}
	return result
}

func GetUserFollowerList(fs []models.Relation) []interface{} {
	result := make([]interface{}, len(fs))
	for i := 0; i < len(fs); i++ {
		result[i] = GetUserFollower(&fs[i])
	}
	return result
}

func GetUserFollowing(f *models.Relation) map[string]interface{} {
	result := map[string]interface{}{
		"id":         f.FromUser.ID,
		"email":      f.FromUser.Email,
		"first_name": f.FromUser.FirstName,
		"last_name":  f.FromUser.LastName,
		"bio":        f.FromUser.Bio,
		"image":      f.FromUser.Email,
	}
	return result
}

func GetUserFollowingList(fs []models.Relation) []interface{} {
	result := make([]interface{}, len(fs))
	for i := 0; i < len(fs); i++ {
		result[i] = GetUserFollowing(&fs[i])
	}
	return result
}
