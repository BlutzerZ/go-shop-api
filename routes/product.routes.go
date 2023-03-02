package routes

import (
	"go-shop-api/controllers"
	"go-shop-api/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {

	authMiddleware, _ := middleware.JWTMiddleware()

	r.GET("/products", authMiddleware.MiddlewareFunc(), controllers.GetProduct)
	r.GET("/product/:id", controllers.GetProductByID)
	r.POST("/product/create", controllers.CreateProduct)
	r.PUT("/product/:id", controllers.ProductUpdate)
	r.DELETE("/product/:id", controllers.ProductDelete)
}
