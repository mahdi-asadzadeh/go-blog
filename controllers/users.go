package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/inputs"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/services"
	"github.com/mahdi-asadzadeh/go-blog/utils"
	"github.com/mahdi-asadzadeh/go-blog/utils/extractors"
	"golang.org/x/crypto/bcrypt"
)

func InitUserRoutes(router *gin.RouterGroup) {
	router.POST("/register", RegisterUser)
	router.POST("/login", UsersLogin)
}

func RegisterUser(ctx *gin.Context) {
	var json inputs.RegisterInput
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	newUser := models.User{
		Password:  string(password),
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Email:     json.Email,
	}
	if err := services.CreateOneUser(&newUser); err != nil {
		utils.APIErrorResponse(ctx, http.StatusUnprocessableEntity, "POST", err.Error())
		return
	}
	utils.APIResponse(
		ctx,
		"User created successfully.",
		http.StatusCreated,
		"POST",
		extractors.GetUser(&newUser),
	)
}

func UsersLogin(ctx *gin.Context) {
	var json inputs.LoginInput
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.APIErrorResponse(ctx, http.StatusBadRequest, "POST", err.Error())
		return
	}
	user, err := services.FindOneUser(&models.User{Email: json.Email})
	if err != nil {
		utils.APIErrorResponse(ctx, http.StatusForbidden, "POST", err.Error())
		return
	}
	if user.IsValidPassword(json.Password) != nil {
		utils.APIErrorResponse(ctx, http.StatusForbidden, "POST", "Invalid credentials")
		return
	}
	utils.APIResponse(
		ctx,
		"User successfully like.",
		http.StatusCreated,
		"POST",
		extractors.CreateLoginSuccessfull(&user),
	)
}
