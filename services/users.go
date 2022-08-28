package services

import (
	"errors"
	"fmt"

	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"golang.org/x/crypto/bcrypt"
)

func SetPassword(user *models.User, password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return nil
}

func FindOneUser(condition interface{}) (models.User, error) {
	fmt.Println(condition)
	database := infrastructure.GetDB()
	var user models.User
	err := database.Where(condition).First(&user).Error
	return user, err
}

func CreateOneUser(data *models.User) error {
	database := infrastructure.GetDB()
	err := database.Create(data).Error
	return err
}
