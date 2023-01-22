package routes

import (
	"go-shop-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	r.POST("/user/auth", controllers.UserAuth)
	r.POST("/user/create", controllers.UserCreate)
	r.DELETE("/user/delete", controllers.UserDelete)
}
