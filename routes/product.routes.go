package routes

import (
	"go-shop-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProduct)
	r.GET("/product/:id", controllers.GetProductByID)
	r.POST("/product/create", controllers.CreateProduct)
	r.PUT("/product/:id", controllers.ProductUpdate)
	r.DELETE("/product/:id", controllers.ProductDelete)
}
