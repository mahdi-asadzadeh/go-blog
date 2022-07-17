package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

func InitTagRoutes(router *gin.RouterGroup) {
	router.GET("/list", TagList)
}

func TagList(ctx *gin.Context) {
	tags, err := services.FetchAllTags()
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfuly tag list.",
		http.StatusOK,
		"GET",
		extractors.GetTagList(tags),
	)
}
