package controllers

import (
	"fmt"
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// func thisIsController() {

// }

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "here the details product"})
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sucess creating username",
		"details": createProduct,
	})
}
