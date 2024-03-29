package extractors

import (
	"github.com/mahdi-asadzadeh/go-blog/models"
)

func GetUser(u *models.User) map[string]interface{} {
	result := map[string]interface{}{
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"bio":        u.Bio,
		"image":      u.Image,
	}
	return result
}

func CreateLoginSuccessfull(user *models.User) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"token":   user.GenerateJwtToken(),
		"user": map[string]interface{}{
			"email": user.Email,
			"id":    user.ID,
		},
	}
}
