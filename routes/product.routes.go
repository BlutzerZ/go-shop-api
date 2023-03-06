package routes

import (
	"go-shop-api/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {

	r.GET("/products", controllers.GetProduct)
	r.GET("/product/:id", controllers.GetProductByID)
	r.POST("/product/create", authMiddleware.MiddlewareFunc(), controllers.CreateProduct)
	r.PUT("/product/:id", authMiddleware.MiddlewareFunc(), controllers.ProductUpdate)
	r.DELETE("/product/:id", authMiddleware.MiddlewareFunc(), controllers.ProductDelete)
}
