package controllers

import (
	"fmt"
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	//
	product, err := configs.GetProductByID(configs.DB, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"result": product})
}

func GetProduct(c *gin.Context) {
	prodLimitQuery := c.Query("limit")

	// change value to int
	prodLimit, err := strconv.Atoi(prodLimitQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// QUERY On Configs
	products, err := configs.GetProductByLimit(configs.DB, prodLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": products})
}

type ProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Desc  string `json:"desc"`
	Stock int    `json:"stock" binding:"required"`
	CatID int    `json:"catid"`
}

func CreateProduct(c *gin.Context) {
	var productRequest ProductRequest

	// Get JSON Body
	err := c.ShouldBindJSON(&productRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)

		}
		return
	}

	// Convert
	createProduct := models.Product{
		Name:       productRequest.Name,
		Desc:       productRequest.Desc,
		Stock:      productRequest.Stock,
		CatID:      productRequest.CatID,
		DateCreate: time.Now().Unix(),
		DateUpdate: time.Now().Unix(),
	}

	// Databse Config
	createProduct, err = configs.AddProduct(configs.DB, createProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sucess creating product",
		"details": createProduct,
	})
}

func ProductDelete(c *gin.Context) {
	productID := c.Param("id")

	// query to db
	err := configs.DeleteProduct(configs.DB, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sucess deleted " + productID})
}
