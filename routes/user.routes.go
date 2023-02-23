package routes

import (
	"go-shop-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.POST("/user/auth", controllers.UserAuth)            //Auth with username and password
	r.POST("/user/create", controllers.UserCreate)        // Create User or SignUp with username and password field
	r.PUT("/user/change", controllers.UserChangePassword) // Change new Password
	r.DELETE("/user/delete", controllers.UserDelete)      // Delete User From database
}
