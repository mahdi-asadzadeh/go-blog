package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/middlewares"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

func InitUserRelationRoutes(router *gin.RouterGroup) {
	router.GET("/followers/:id", ListFollowers)
	router.GET("/following/:id", ListFollowing)
	router.Use(middlewares.RequiredAuthMiddleware())
	{
		router.POST("/follow/:id", FollowUser)
		router.DELETE("/unfollow/:id", UnFollowUser)
	}

}

func ListFollowers(ctx *gin.Context) {
	userIDStr, _ := strconv.Atoi(ctx.Param("id"))
	userID := uint(userIDStr)
	database := infrastructure.GetDB()

	var relations []models.Relation
	database.Where(&models.Relation{FromUserID: userID}).Preload("ToUser").Find(&relations)
	utils.APIResponse(ctx, "", http.StatusOK, "GET", extractors.GetUserFollowerList(relations))
}

func ListFollowing(ctx *gin.Context) {
	userIDStr, _ := strconv.Atoi(ctx.Param("id"))
	userID := uint(userIDStr)
	database := infrastructure.GetDB()

	var relations []models.Relation
	database.Where(&models.Relation{ToUserID: userID}).Preload("FromUser").Find(&relations)
	utils.APIResponse(ctx, "", http.StatusOK, "GET", extractors.GetUserFollowingList(relations))
}

func FollowUser(ctx *gin.Context) {
	fromUser := ctx.MustGet("currentUser").(models.User)
	toUserID, _ := strconv.Atoi(ctx.Param("id"))
	database := infrastructure.GetDB()

	var relation models.Relation
	database.Where(&models.Relation{FromUserID: fromUser.ID, ToUserID: uint(toUserID)}).
		Attrs(models.Relation{FromUserID: fromUser.ID, ToUserID: uint(toUserID)}).
		FirstOrCreate(&relation)
	utils.APIResponse(ctx, "", http.StatusOK, "POST", extractors.GetUserFollower(&relation))
}

func UnFollowUser(ctx *gin.Context) {
	fromUser := ctx.MustGet("currentUser").(models.User)
	toUserID, _ := strconv.Atoi(ctx.Param("id"))
	database := infrastructure.GetDB()
	relation := models.Relation{FromUserID: fromUser.ID, ToUserID: uint(toUserID)}

	database.Where(&relation).Delete(&relation)
	utils.APIResponse(ctx, "", http.StatusOK, "DELETE", extractors.GetUserFollowing(&relation))
}
