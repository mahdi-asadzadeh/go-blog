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

// @Summary Register an user
// @Description Register an user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body inputs.RegisterInput true "Register an user"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 422 {object} utils.ErrorResponse
// @Router /user/register [Post]
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

// @Summary Login an user
// @Description Login an user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body inputs.LoginInput true "Login an user"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Router /user/login [Post]
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
