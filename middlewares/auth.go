package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-asadzadeh/go-blog/infrastructure"
	"github.com/mahdi-asadzadeh/go-blog/models"
	"github.com/mahdi-asadzadeh/go-blog/utils"
)

func RequiredAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if exists && user.(models.User).ID != 0 {
			return
		} else {
			err, exists := c.Get("authErr")
			if exists {
				c.AbortWithStatusJSON(http.StatusForbidden, utils.CreateDetailedError("auth_error", err.(error)))
			} else {
				c.JSON(http.StatusForbidden, utils.CreateErrorWithMessage("You must be authenticated"))
				c.Abort()
			}
		}
	}
}

func UserLoadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			jwtParts := strings.Split(bearer, " ")
			if len(jwtParts) == 2 {
				jwtEncoded := jwtParts[1]

				token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
					}
					secret := []byte(os.Getenv("JWT_SECRET"))
					return secret, nil
				})

				if err != nil {
					println(err.Error())
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					if userId, ok := claims["user_id"]; ok {
						userId = uint(userId.(float64))
						fmt.Printf("[+] Authenticated request, authenticated user id is %d\n", userId)

						var user models.User
						if userId != 0 {
							database := infrastructure.GetDB()
							database.Preload("Roles").First(&user, userId)
							c.Set("currentUser", user)
							c.Set("currentUserId", user.ID)
						}
					}
				} else {

				}
			}
		}
	}
}
