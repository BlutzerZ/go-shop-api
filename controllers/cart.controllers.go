package controllers

import (
	"fmt"
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type CartItemRequest struct {
	Qty int `json:"qty" binding:"required"`
}

func AddItemToCart(c *gin.Context) {
	var item models.CartItem
	var cartItemReq CartItemRequest

	// GET PHARAM OF URL
	productID := c.Param("productID")

	// GET JSON BODY
	err := c.ShouldBindJSON(&cartItemReq)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
		}
		return
	}

	// CLAIM JWT TOKEN
	claims := jwt.ExtractClaims(c)
	userID := uuid.FromStringOrNil(claims["uuid"].(string))

	item.ProductID = uuid.FromStringOrNil(productID) // this is will be good if i modified product id as uuid type

	err = configs.InsertItemToCart(configs.DB, userID, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}

func EditItemOnCart(c *gin.Context) {

}

func DeleteItemOnCart(c *gin.Context) {

}

func GetItemsOnCart(c *gin.Context) {

}

func SendCartToCheckout(c *gin.Context) {

}
