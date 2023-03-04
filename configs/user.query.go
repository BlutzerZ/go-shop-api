package configs

import (
	"fmt"
	"go-shop-api/models"

	"gorm.io/gorm"
)

// ================
//
//	USER QUERY
//
// ================
func AuthUser(db *gorm.DB, username string, password string) (bool, error) {
	var user models.User

	err := db.Find(&user, "username = ? AND password = ?", username, password).Error
	isAuth := false
	if user != (models.User{}) {
		isAuth = true
	}

	return isAuth, err
}

func AddUser(db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func DeleteUser(db *gorm.DB, username string) (string, error) {
	err := db.Delete(models.User{}, "username = ?", username).Error

	return username, err
}

type UserChangePasswordRequest struct {
	Username        string
	CurrentPassword string `json:"currpwd" binding:"required"`
	NewPassword     string `json:"newpwd" binding:"required"`
}

func ChangePasswordUser(db *gorm.DB, UpdatedUser UserChangePasswordRequest) (bool, error) {
	var user models.User

	err := db.Find(&user, "username = ? AND password = ?", UpdatedUser.Username, UpdatedUser.CurrentPassword).Error
	isChanged := false
	if user != (models.User{}) {
		fmt.Println(user)
		err = db.Model(&user).Where("username = ? AND password = ?", UpdatedUser.Username, UpdatedUser.CurrentPassword).Update("password", UpdatedUser.NewPassword).Error
		isChanged = true
	}

	return isChanged, err
}
