package controllers

import (
	"fmt"
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"
	"time"

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

	// forward request to item cart model
	item.ProductID = uuid.FromStringOrNil(productID) // this is will be good if i modified product id as uuid type
	item.Qty = cartItemReq.Qty
	item.DateUpdate = time.Now().Unix()

	err = configs.InsertItemToCart(configs.DB, userID, &item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"details": item,
		"message": "success create item",
	})
}

func EditItemOnCart(c *gin.Context) {
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

	// forward request to item cart model
	item.ProductID = uuid.FromStringOrNil(productID) // this is will be good if i modified product id as uuid type
	item.Qty = cartItemReq.Qty
	item.DateUpdate = time.Now().Unix()

	// QUERY
	err = configs.EditItemCart(configs.DB, userID, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"edited":  cartItemReq,
		"message": "success edit item",
	})
}

type DeleteCartItemConfirmationRequest struct {
	Confirmation string `json:"confirmation" bimding:"required"`
}

func DeleteItemOnCart(c *gin.Context) {
	var item models.CartItem
	var deleteCartReq DeleteCartItemConfirmationRequest

	// GET PHARAM OF URL
	productID := c.Param("productID")

	// GET JSON BODY
	err := c.ShouldBindJSON(&deleteCartReq)
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

	// forward request to item cart model
	item.ProductID = uuid.FromStringOrNil(productID) // this is will be good if i modified product id as uuid type

	// QUERY
	err = configs.DeleteItemCart(configs.DB, userID, item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product-id": item.ProductID,
		"message":    "success delete item",
	})

}

func GetItemsOnCart(c *gin.Context) {
	var items []models.CartItem

	// CLAIM JWT TOKEN
	claims := jwt.ExtractClaims(c)
	userID := uuid.FromStringOrNil(claims["uuid"].(string))

	items, err := configs.GetAllItemCart(configs.DB, userID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item":    items,
		"message": "success generate result",
	})

}

func SendCartToOrder(c *gin.Context) {
	var itemsOfCart []models.CartItem

	// CLAIM JWT TOKEN
	claims := jwt.ExtractClaims(c)
	userID := uuid.FromStringOrNil(claims["uuid"].(string))

	// query
	itemsOfCart, err := configs.GetAllItemCart(configs.DB, userID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	// send item to order
	err = configs.ForwardCartToOrder(configs.DB, userID, itemsOfCart)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	// delete all item cart
	err = configs.DeleteAllItemCart(configs.DB, userID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sucess added to order",
	})
}
