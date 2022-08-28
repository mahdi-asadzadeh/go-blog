package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

// @Summary List an user's followers
// @Description List an user's followers
// @Tags Relation
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Router /relation/followers{id} [GET]
func ListFollowers(ctx *gin.Context) {
	userIDStr, _ := strconv.Atoi(ctx.Param("id"))
	userID := uint(userIDStr)
	database := infrastructure.GetDB()

	var relations []models.Relation
	database.Where(&models.Relation{FromUserID: userID}).Preload("ToUser").Find(&relations)
	utils.APIResponse(ctx, "", http.StatusOK, "GET", extractors.GetUserFollowerList(relations))
}

// @Summary List an user's followings
// @Description  List an user's followings
// @Tags Relation
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Router /relation/followings{id} [GET]
func ListFollowing(ctx *gin.Context) {
	userIDStr, _ := strconv.Atoi(ctx.Param("id"))
	userID := uint(userIDStr)
	database := infrastructure.GetDB()

	var relations []models.Relation
	database.Where(&models.Relation{ToUserID: userID}).Preload("FromUser").Find(&relations)
	utils.APIResponse(ctx, "", http.StatusOK, "GET", extractors.GetUserFollowingList(relations))
}

// @Security Authorization
// @Summary Follow an user
// @Description Follow an user
// @Tags Relation
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Router /relation/follow/{id} [POST]
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

// @Security Authorization
// @Summary Unfollow an user
// @Description Unfollow an user
// @Tags Relation
// @Accept json
// @Product json
// @Param id path string true "id"
// @Success 200 {object} utils.Response
// @Router /relation/unfollow/{id} [DELETE]
func UnFollowUser(ctx *gin.Context) {
	fromUser := ctx.MustGet("currentUser").(models.User)
	toUserID, _ := strconv.Atoi(ctx.Param("id"))
	database := infrastructure.GetDB()
	relation := models.Relation{FromUserID: fromUser.ID, ToUserID: uint(toUserID)}

	database.Where(&relation).Delete(&relation)
	utils.APIResponse(ctx, "", http.StatusOK, "DELETE", extractors.GetUserFollowing(&relation))
}
