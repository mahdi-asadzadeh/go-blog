package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/inputs"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

// @Security Authorization
// @Summart Create a comment by article's slug
// @Description Create a comment by article's slug
// @Tags Comment
// @Accept json
// @Product json
// @Param request body inputs.CreateCommentInput true "Create comment"
// @Param slug path string true "slug"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 422 {object} utils.ErrorResponse
// @Router /comment/create/{slug} [POST]
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

// @Summart List comments by article's slug
// @Description List comments by article's slug
// @Tags Comment
// @Accept json
// @Product json
// @Param slug path string true "slug"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /comment/list/{slug} [GET]
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

// @Summart Show a comment
// @Description Show a comment
// @Tags Comment
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Router /comment/show/{id} [GET]
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

// @Security Authorization
// @Summart Delete a comment
// @Description Delete a comment
// @Tags Comment
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Router /comment/delete/{id} [DELETE]
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
