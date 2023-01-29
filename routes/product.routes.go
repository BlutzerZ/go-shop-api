package routes

import (
	"go-shop-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	// r.GET("/product/", controllers.GetProduct)
	// r.GET("/product/:id", controllers.GetProductByID)
	r.POST("/product/create", controllers.CreateProduct)
	// r.GET("/product/delete", controllers.ProductDelete)
}
