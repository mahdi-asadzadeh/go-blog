package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string
	Description string
	Users       []User `gorm:"many2many:users_roles;"`
}

type UserRole struct {
	User   User `gorm:"foreignkey:UserId"`
	UserId uint
	Role   User `gorm:"foreignkey:RoleId"`
	RoleId uint
}

func (UserRole) TableName() string {
	return "users_roles"
}
