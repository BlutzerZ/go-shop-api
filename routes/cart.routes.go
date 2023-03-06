package routes

import (
	"go-shop-api/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/cart/:productID", authMiddleware.MiddlewareFunc(), controllers.AddItemToCart)
	// r.PUT("/cart/:productID", controller.EditItemOnCart)
	// r.DELETE("/cart/:productID", controllers.DeleteItemOnCart)
	// r.GET("/cart", controllers.GetItemsOnCart)

	// r.POST("/cart/checkout", controller.SendCartToCheckout)
}
