package routes

import (
	"go-shop-api/controllers"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.GET("/order", authMiddleware.MiddlewareFunc(), controllers.GetOrders)
	r.GET("/order/:orderID", authMiddleware.MiddlewareFunc(), controllers.GetOrderByID)
	r.PUT("/order/:orderID/cancel", authMiddleware.MiddlewareFunc(), controllers.CancelOrderByID)
	r.DELETE("/order/:orderID", authMiddleware.MiddlewareFunc(), controllers.DeleteOrderByID)
}
