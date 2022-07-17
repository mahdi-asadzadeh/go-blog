package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Openiing a database and save the reference to `Database` struct.
func InitDB() *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}
	DB = db
	return db
}

// Using this fuction to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
