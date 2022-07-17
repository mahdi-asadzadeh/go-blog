package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/middlewares"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

func InitLikeRoutes(router *gin.RouterGroup) {
	router.GET("/likes/:slug", LikesArticle)
	router.Use(middlewares.RequiredAuthMiddleware())
	{
		router.POST("/create/:slug", LikeArticle)
		router.DELETE("/delete/:slug", DisLikeArticle)
		router.GET("/my", MyLikes)
	}
}

func MyLikes(ctx *gin.Context) {
	database := infrastructure.GetDB()

	var likes []models.Like
	user_id := ctx.MustGet("currentUserId")
	err := database.Model(&likes).Where("user_id = ?", user_id).Preload("User").Find(&likes).Error
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfuly your like list.",
		http.StatusOK,
		"GET",
		extractors.GetLikeList(likes),
	)
}

func LikesArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	database := infrastructure.GetDB()

	var article models.Article
	err := database.Model(&article).Where("slug = ?", slug).Select([]string{"id"}).First(&article).Error
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}

	var likes []models.Like
	err = database.Model(&likes).Where("article_id = ?", article.ID).Preload("User").Preload("Article").Find(&likes).Error
	utils.APIResponse(
		ctx,
		"Successfuly article like list.",
		http.StatusOK,
		"GET",
		extractors.GetLikeList(likes),
	)
}

func DisLikeArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	database := infrastructure.GetDB()
	var result struct {
		ID string
	}
	err := database.Table("articles").Select("id").Where("slug = ?", slug).Scan(&result).Error
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	user := ctx.MustGet("currentUser").(models.User)
	var like models.Like
	err = database.Model(models.Like{}).Where("user_id = ? AND article_id = ?", user.ID, result.ID).First(&like).Error

	if err == nil {
		database.Delete(&like)
		utils.APIResponse(
			ctx,
			"Successfuly dislike article.",
			http.StatusOK,
			"DELETE",
			nil,
		)
		return
	} else {
		utils.APIErrorResponse(ctx, http.StatusForbidden, "DELETE", "You were not liking this article, so you can not perform this operation")
	}
}

func LikeArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")

	database := infrastructure.GetDB()
	var article models.Article
	err := database.Model(&article).Where("slug = ?", slug).Select([]string{"id", "title"}).First(&article).Error

	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "POST", err.Error())
		return
	}
	user := ctx.MustGet("currentUser").(models.User)
	alreadyLike := services.IsLikeBy(&article, &user)
	if !alreadyLike {
		like := models.Like{
			ArticleID: article.ID,
			UserID:    user.ID,
		}
		err := services.CreateOneLike(&like)
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusUnprocessableEntity, "POST", err.Error())
			return
		}
		utils.APIResponse(
			ctx,
			fmt.Sprintf("You liked the article \"%v\" successfully", article.Title),
			http.StatusOK,
			"POST",
			nil,
		)
	} else {
		utils.APIResponse(
			ctx,
			"You have already liked this article.",
			http.StatusForbidden,
			"POST",
			nil,
		)
	}
}
