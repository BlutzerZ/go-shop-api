package controllers

import (
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// ===============================================
//
// 		Q U E R Y      T O     O R D E R
//
// ===============================================

func GetOrders(c *gin.Context) {
	var orders []models.Order

	claims := jwt.ExtractClaims(c)
	userID := uuid.FromStringOrNil(claims["uuid"].(string))

	orders, err := configs.GetAllOrder(configs.DB, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":  orders,
		"message": "success retrive orders",
	})

}

func GetOrderByID(c *gin.Context) {
	var orderItems []models.OrderItem

	orderID := uuid.FromStringOrNil(c.Param("orderID"))

	claims := jwt.ExtractClaims(c)
	userID := uuid.FromStringOrNil(claims["uuid"].(string))

	orderItems, err := configs.GetOrderByID(configs.DB, userID, orderID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": orderItems,
		"message": "success retrive order items",
	})
}

func CancelOrderByID(c *gin.Context) {
	var order models.Order

	orderIDParam := c.Param("orderID")
	order.ID = uuid.FromStringOrNil(orderIDParam)
	claims := jwt.ExtractClaims(c)
	order.UserID = uuid.FromStringOrNil(claims["uuid"].(string))

	err := configs.CancelOrderByID(configs.DB, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed canceling transaction",
			"error":   err,
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "sucess canceling transaction id " + c.Param("orderID"),
	})
}

func DeleteOrderByID(c *gin.Context) {
	var order models.Order

	order.ID = uuid.FromStringOrNil(c.Param("orderID"))
	claims := jwt.ExtractClaims(c)
	order.UserID = uuid.FromStringOrNil(claims["uuid"].(string))

	err := configs.DeleteAllOrderItem(configs.DB, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete order",
			"error":   err,
		})
		return
	}

	err = configs.DeleteOrderByID(configs.DB, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete order",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success to delete order",
	})
}
