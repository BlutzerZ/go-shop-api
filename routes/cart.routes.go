package routes

import (
	"go-shop-api/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/cart/:productID", authMiddleware.MiddlewareFunc(), controllers.AddItemToCart)
	r.PUT("/cart/:productID", authMiddleware.MiddlewareFunc(), controllers.EditItemOnCart)
	r.DELETE("/cart/:productID", authMiddleware.MiddlewareFunc(), controllers.DeleteItemOnCart)
	r.GET("/cart", authMiddleware.MiddlewareFunc(), controllers.GetItemsOnCart)

	// r.POST("/cart/checkout", controller.SendCartToCheckout)
}
