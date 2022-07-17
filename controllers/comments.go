package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/inputs"
	"github.com/mahdi-asadzadeh/go-blog/middlewares"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

func InitCommentRoutes(router *gin.RouterGroup) {
	router.GET("/list/:slug", ListComments)
	router.GET("/show/:id", ShowComment)

	router.Use(middlewares.RequiredAuthMiddleware())
	{
		router.POST("/create/:slug", CreateComment)
		router.DELETE("/delete/:id", DeleteComment)
	}
}

func CreateComment(ctx *gin.Context) {
	slug := ctx.Param("slug")
	articleId, err := services.FetchArticleId(slug)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "POST", err.Error())
		return
	}
	var json inputs.CreateCommentInput
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}

	comment := models.Comment{
		Content:   json.Content,
		ArticleID: articleId,
		User:      ctx.MustGet("currentUser").(models.User),
	}

	if err = services.CreateOne(&comment); err != nil {
		utils.APIErrorResponse(ctx, http.StatusUnprocessableEntity, "POST", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfully create a comment.",
		http.StatusOK,
		"POST",
		extractors.GetComment(&comment),
	)
}

func ListComments(ctx *gin.Context) {
	slug := ctx.Param("slug")

	articleId, err := services.FetchArticleId(slug)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "POST", err.Error())
		return
	}

	var comments []models.Comment
	database := infrastructure.GetDB()

	err = database.Model(&comments).Where("article_id = ?", articleId).Find(&comments).Error
	ctx.JSON(http.StatusOK, comments)
	utils.APIResponse(
		ctx,
		"Successfully comment list.",
		http.StatusOK,
		"GET",
		extractors.GetCommentList(comments),
	)
}

func ShowComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	id, _ := strconv.Atoi(commentID)
	comment, err := services.FetchCommentByID(id)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "GET", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfully get a comment.",
		http.StatusOK,
		"GET",
		extractors.GetComment(&comment),
	)
}

func DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("id")
	id, _ := strconv.Atoi(commentID)
	comment, err := services.FetchCommentByID(id)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "DELETE", err.Error())
		return
	}
	user := ctx.MustGet("currentUser").(models.User)

	database := infrastructure.GetDB()
	if comment.UserID == user.ID {
		err := database.Delete(&comment).Error
		if err != nil {
			utils.APIErrorResponse(ctx, http.StatusBadRequest, "DELETE", err.Error())
			return
		}
		utils.APIResponse(ctx, "Comment Deleted successfully", http.StatusOK, "DELETE", nil)
		return
	} else {
		utils.APIErrorResponse(
			ctx,
			http.StatusForbidden,
			"DELETE",
			"You have to be admin or the owner of this comment to delete it.",
		)
		return
	}

}
