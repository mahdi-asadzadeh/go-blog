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


	docs "github.com/mahdi-asadzadeh/go-blog/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
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

	docs.SwaggerInfo.BasePath = "/api"


	routerAPI := router.Group("/api")
	{
		userG := routerAPI.Group("user")
		{
			userG.POST("/register", controllers.RegisterUser)
			userG.POST("/login", controllers.UsersLogin)
		}

		articleG := routerAPI.Group("article")
		{
			articleG.GET("/list", controllers.ListArticles)
			articleG.GET("/detail/:slug", controllers.DetailArticle)
			articleG.GET("/images/:slug", controllers.ArticleImageList)
			articleG.Use(middlewares.RequiredAuthMiddleware())
			{
				articleG.POST("/create", controllers.CreateArticle)
				articleG.POST("/upload-images/:slug", controllers.UploadImagesProduct)
				articleG.DELETE("/delete/:slug", controllers.DeleteArticle)
			}
		}

		likeG := routerAPI.Group("likes")
		{
			likeG.GET("/likes/:slug", controllers.LikesArticle)
			likeG.Use(middlewares.RequiredAuthMiddleware())
			{
				likeG.POST("/create/:slug", controllers.LikeArticle)
				likeG.DELETE("/delete/:slug", controllers.DisLikeArticle)
				likeG.GET("/my", controllers.MyLikes)
			}
		}

		categoryG := routerAPI.Group("category")
		{
			categoryG.GET("/list", controllers.ListCategories)
			categoryG.GET("/detail/:slug", controllers.GategoryArticles)
		}

		commentG := routerAPI.Group("comment")
		{
			commentG.GET("/list/:slug", controllers.ListComments)
			commentG.GET("/show/:id", controllers.ShowComment)
		
			commentG.Use(middlewares.RequiredAuthMiddleware())
			{
				commentG.POST("/create/:slug", controllers.CreateComment)
				commentG.DELETE("/delete/:id", controllers.DeleteComment)
			}
		}

		tagG := routerAPI.Group("tag")
		{
			tagG.GET("/list", controllers.TagList)
		}

		relationG := routerAPI.Group("relation")
		{
			relationG.GET("/followers/:id", controllers.ListFollowers)
			relationG.GET("/followings/:id", controllers.ListFollowing)
			relationG.Use(middlewares.RequiredAuthMiddleware())
			{
				relationG.POST("/follow/:id", controllers.FollowUser)
				relationG.DELETE("/unfollow/:id", controllers.UnFollowUser)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.StaticFS("/file", http.Dir("public")) // Serve static files
	router.Run(":8080")
}
