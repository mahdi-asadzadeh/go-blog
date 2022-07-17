package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/inputs"
	"github.com/mahdi-asadzadeh/go-blog/middlewares"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
	"gorm.io/gorm"
)

func InitArticleRoutes(router *gin.RouterGroup) {
	router.GET("/list", ListArticles)
	router.GET("/detail/:slug", DetailArticle)
	router.GET("/images/:slug", ArticleImageList)
	router.Use(middlewares.RequiredAuthMiddleware())
	{
		router.POST("/create", CreateArticle)
		router.POST("/upload-images/:slug", UploadImagesProduct)
		router.DELETE("/delete/:slug", DeleteArticle)
	}

}

func UploadImagesProduct(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	slug := ctx.Param("slug")

	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}

	articleID, err := services.FetchArticleId(slug)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "POST", err.Error())
		return
	}

	newFileName := uuid.New().String() + file.Filename
	err = ctx.SaveUploadedFile(file, "public/"+newFileName)

	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	err = services.CreateOneImageArticle(&models.ArticleImage{
		Path:      newFileName,
		Size:      file.Size,
		ArticleID: articleID,
	})
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly upload image.", http.StatusOK, "POST", nil)
}

func ArticleImageList(ctx *gin.Context) {
	slug := ctx.Param("slug")
	articleID, err := services.FetchArticleId(slug)
	if err != nil {
		utils.APIResponse(ctx, "Article not found.", http.StatusNotFound, "GET", err.Error())
		return
	}
	images, err := services.FetchArticleImages(articleID)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "GET", err.Error())
		return
	}
	utils.APIResponse(ctx, "Successfuly list images.", http.StatusOK, "GET", extractors.GetArticleImageList(images))

}

func DetailArticle(ctx *gin.Context) {
	slugParam := ctx.Param("slug")
	article, err := services.FetchArticleDetails(&models.Article{Slug: slugParam})

	if err == gorm.ErrRecordNotFound {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "GET", err.Error())
		return
	} else if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "GET", err.Error())
	}

	utils.APIResponse(
		ctx,
		"Successfuly article detail.",
		http.StatusOK,
		"GET",
		extractors.GetArticleDetail(&article),
	)
}

func ListArticles(ctx *gin.Context) {
	pageSizeStr := ctx.Query("page_size")
	pageStr := ctx.Query("page")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 5
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	articles, modelsCount, _ := services.FetchArticlesPage(page, pageSize)
	utils.APIResponse(
		ctx,
		"Successfuly article list.",
		http.StatusOK,
		"GET",
		extractors.GetArticleListPage(ctx.Request, articles, page, pageSize, modelsCount),
	)
}

func CreateArticle(ctx *gin.Context) {
	var json inputs.CreateArticleInput
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}

	database := infrastructure.GetDB()
	categories := make([]models.Category, len(json.Categories))
	tags := make([]models.Tag, len(json.Tags))

	for index, tag := range json.Tags {
		database.Where(&models.Tag{Slug: slug.Make(tag.Name)}).
			Attrs(models.Tag{Name: tag.Name, Description: tag.Description}).
			FirstOrCreate(&tags[index])
	}

	for index, _ := range json.Categories {
		database.Where(&models.Category{Slug: slug.Make(json.Categories[index].Name)}).
			Attrs(models.Category{Name: json.Categories[index].Name, Description: json.Categories[index].Description}).
			FirstOrCreate(&categories[index])
	}

	newArticle := models.Article{
		Title:       json.Title,
		Description: json.Description,
		Body:        json.Body,
		UserID:      ctx.MustGet("currentUserId").(uint),
		User:        ctx.MustGet("currentUser").(models.User),
		Categories:  categories,
		Tags:        tags,
	}
	if err := services.CreateOneArticle(&newArticle); err != nil {
		utils.APIErrorResponse(ctx, http.StatusUnprocessableEntity, "POST", err.Error())
		return
	}

	utils.APIResponse(
		ctx,
		"Successfuly create article.",
		http.StatusOK,
		"POST",
		extractors.GetArticleDetail(&newArticle),
	)
}

func DeleteArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")

	user := ctx.MustGet("currentUser").(models.User)
	err := services.DeleteArticleIfOwnerOrAdmin(&user, &models.Article{Slug: slug})
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusNotFound, "DELETEdd", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"Article deleted successfully.",
		http.StatusOK,
		"DELETE",
		nil,
	)
}
