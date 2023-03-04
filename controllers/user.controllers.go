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
)

type UserAuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserAuth(c *gin.Context) {
	var authInput UserAuthRequest

	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		fmt.Println(err)
	}

	// auth user and password
	isAuth, err := configs.AuthUser(configs.DB, authInput.Username, authInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// return json is auth
	if isAuth {
		c.JSON(http.StatusOK, gin.H{"message": "login sucess"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login fail"})
	}

}

type UserRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserCreate(c *gin.Context) {
	var userRequest UserRequest

	// Get Json Body
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)

		}
		return
	}

	// Convert
	createUser := models.User{
		Email:      userRequest.Email,
		Username:   userRequest.Username,
		Password:   userRequest.Password,
		DateCreate: time.Now().Unix(),
		DateUpdate: time.Now().Unix(),
	}

	//Database Config
	createUser, err = configs.AddUser(configs.DB, createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success creating username",
		"details": createUser,
	})
}

type UserDeleteRequest struct {
	Confirmation string `json:"confirmation" binding:"required"`
}

func UserDelete(c *gin.Context) {
	var ud UserDeleteRequest

	// EXTRACT JSON BODY
	err := c.ShouldBindJSON(&ud)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)

		}
		return
	}

	// CLAIM TOKEN IF JSON BODY SAY "YES", THEN USE RESULT TO QUERY DELETE
	if ud.Confirmation == "yes" {
		claims := jwt.ExtractClaims(c)
		username := claims["Username"].(string)
		configs.DeleteUser(configs.DB, username)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s sucess deleted from database", username)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "error not confirmed"})
}

// type UserChangePasswordRequest struct {
// 	Username        string
// 	currentPassword int `json:"currpwd" binding:"required"`
// 	newPassword     int `json:"newpwd" binding:"required"`
// }

func UserChangePassword(c *gin.Context) {
	var changePwReq configs.UserChangePasswordRequest

	err := c.ShouldBindJSON(&changePwReq)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error in field %s condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)

		}
		return
	}
	fmt.Println(changePwReq)

	claims := jwt.ExtractClaims(c)
	username := claims["Username"].(string)
	changePwReq.Username = username

	isChanged, err := configs.ChangePasswordUser(configs.DB, changePwReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if isChanged {
		c.JSON(http.StatusOK, gin.H{
			"message": "success change password",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to change password",
		})
	}
}
