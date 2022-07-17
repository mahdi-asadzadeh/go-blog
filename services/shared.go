package services

import (
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
)

func CreateOne(data interface{}) error {
	database := infrastructure.GetDB()
	err := database.Save(data).Error
	return err
}
