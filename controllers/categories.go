package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

// @Summary Category's list articles
// @Description Category's list articles
// @Tags Category
// @Accept json
// @Product json
// @Param slug path string true "Article's slug"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /category/detail/{slug} [GET]
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

// @Summary List categories
// @Description List categories
// @Tags Category
// @Accept json
// @Product json
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /category/list [GET]
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
