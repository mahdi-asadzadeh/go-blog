package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mahdi-asadzadeh/go-blog/controllers"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/middlewares"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Article{})
	db.Debug().AutoMigrate(&models.Like{})
	db.Debug().AutoMigrate(&models.Category{})
	db.Debug().AutoMigrate(&models.Comment{})
	db.Debug().AutoMigrate(&models.Tag{})
	db.Debug().AutoMigrate(&models.ArticleTag{})
	db.Debug().AutoMigrate(&models.Relation{})
	db.Debug().AutoMigrate(&models.ArticleImage{})
}

func create(db *gorm.DB) {
	migrate(db)
}

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	database := infrastructure.InitDB()

	args := os.Args
	if len(args) > 1 {
		first := args[1]
		if first == "create" {
			create(database)
		} else if first == "migrate" {
			migrate(database)
		}
	}

	router := gin.Default()
	router.Use(middlewares.UserLoadMiddleware())

	apiRouteGroup := router.Group("/api")

	controllers.InitUserRoutes(apiRouteGroup)
	controllers.InitArticleRoutes(apiRouteGroup.Group("articles"))
	controllers.InitLikeRoutes(apiRouteGroup.Group("likes"))
	controllers.InitCategoryRoutes(apiRouteGroup.Group("categories"))
	controllers.InitCommentRoutes(apiRouteGroup.Group("comments"))
	controllers.InitTagRoutes(apiRouteGroup.Group("tags"))
	controllers.InitUserRelationRoutes(apiRouteGroup.Group("relation"))

	router.StaticFS("/file", http.Dir("public")) // Serve static files
	router.Run(":8080")
}
