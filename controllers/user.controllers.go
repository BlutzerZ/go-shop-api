package controllers

import (
	"fmt"
	"go-shop-api/configs"
	"go-shop-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthUser struct {
	Username string
	Password string
}

func UserAuth(c *gin.Context) {
	var authInput AuthUser

	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		fmt.Println(err)
	}

	// auth user and password
	if authInput.Username == "admin" && authInput.Password == "admin" {
		c.JSON(http.StatusOK, gin.H{"message": "login sucess"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "login failed"})
}

type UserRequest struct {
	Email    string `json:"email" binding:"require"`
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
}

func UserCreate(c *gin.Context) {
	var userRequest UserRequest

	// Convert
	createUser := models.User{
		Email:      userRequest.Email,
		Username:   userRequest.Username,
		Password:   userRequest.Password,
		DateCreate: time.Now().Unix(),
		DateUpdate: time.Now().Unix(),
	}

	// Get Json Body
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		fmt.Println(err)
	}

	//Database Config
	createUser, err = configs.AddUser(configs.DB, createUser)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success creating username",
		"details": createUser})
}

func UserDelete(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{"message": id + " Success Deleted"})
}
