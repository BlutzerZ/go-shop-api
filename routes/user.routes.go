package routes

import (
	"go-shop-api/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/user/auth", authMiddleware.LoginHandler)                                      //Auth with username and password
	r.POST("/user/create", controllers.UserCreate)                                         // Create User or SignUp with username and password field
	r.PUT("/user/change", authMiddleware.MiddlewareFunc(), controllers.UserChangePassword) // Change new Password
	r.DELETE("/user/delete", authMiddleware.MiddlewareFunc(), controllers.UserDelete)      // Delete User From database
}
