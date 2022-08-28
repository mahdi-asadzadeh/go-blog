package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
)

// @Summary List tags
// @Description List tags
// @Tags Tag
// @Accept json
// @Product json
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.ErrorResponse
// @Router /tag/list [GET]
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
