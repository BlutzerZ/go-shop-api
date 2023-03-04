package routes

import (
	"go-shop-api/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	r.POST("/cart/:productID", controllers.AddItemToCart)
	// r.PUT("/cart/:productID", controller.EditItemOnCart)
	// r.DELETE("/cart/:productID", controllers.DeleteItemOnCart)
	// r.GET("/cart", controllers.GetItemsOnCart)

	// r.POST("/cart/checkout", controller.SendCartToCheckout)
}
