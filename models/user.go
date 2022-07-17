package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"column:email;UNIQUE; not null"`
	FirstName string `gorm:"varchar(255);not null"`
	LastName  string `gorm:"varchar(255);not null"`
	Bio       string `gorm:"column:bio;size:1024"`
	Image     string `gorm:"column:image"`
	Password  string `gorm:"column:password;not null"`

	Roles []Role `gorm:"many2many:users_roles"`
}

func (user *User) BeforeSave(db *gorm.DB) (err error) {
	if len(user.Roles) == 0 {
		role := Role{}
		userRole := Role{}
		db.Model(&role).Where("name = ?", "ROLE_USER").First(userRole)
		user.Roles = append(user.Roles, userRole)
	}
	return
}

func (user *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashPassword := []byte(user.Password)
	return bcrypt.CompareHashAndPassword(byteHashPassword, bytePassword)
}

func (user *User) GenerateJwtToken() string {
	jwt_token := jwt.New(jwt.SigningMethodHS512)
	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)

	}

	jwt_token.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   roles,
		"exp":     time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
