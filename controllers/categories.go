package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

func InitCategoryRoutes(router *gin.RouterGroup) {
	router.GET("/list", ListCategories)
	router.GET("/detail/:slug", GategoryArticles)
}

func GategoryArticles(ctx *gin.Context) {
	slug := ctx.Param("slug")
	cat, err := services.FetchCategoryArticles(slug)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfuly detail category.",
		http.StatusOK,
		"GET",
		extractors.GetArticleList(cat.Articles),
	)
}

func ListCategories(ctx *gin.Context) {
	categories, err := services.FetchAllCategories()
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Successfuly category list.",
		http.StatusOK,
		"GET",
		extractors.GetCategoryList(categories),
	)
}
