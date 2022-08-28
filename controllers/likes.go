package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

// @Security Authorization
// @Summary List my likes
// @Description List my likes
// @Tags Like
// @Access json
// @Product json
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /like/my [GET]
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

// @Summary List an article's likes
// @Description List an article's likes
// @Tags Like
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /like/likes/{slug} [GET]
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

// @Security Authorization
// @Summary Dislike an article
// @Description Dislike an article
// @Tags Like
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Router /like/delete/{slug} [DELETE]
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

// @Security Authorization
// @Summary Like an articlle
// @Description Like an article
// @Tags Like
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} utils.Response
// @Failure 403 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 422 {object} utils.ErrorResponse
// @Router /like/create/{slug} [POST]
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
