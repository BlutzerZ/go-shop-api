package configs

import (
	"fmt"
	"go-shop-api/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ================
//
//	USER QUERY
//
// ================
func AuthUser(db *gorm.DB, username string, password string) (bool, uuid.UUID, error) {
	var user models.User

	err := db.Find(&user, "username = ? AND password = ?", username, password).Error
	isAuth := false
	if user != (models.User{}) {
		isAuth = true
	}

	return isAuth, user.UUID, err
}

func AddUser(db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func DeleteUser(db *gorm.DB, uuid uuid.UUID) (uuid.UUID, error) {
	err := db.Delete(models.User{}, "uuid = ?", uuid).Error

	return uuid, err
}

type UserChangePasswordRequest struct {
	UUID            uuid.UUID
	CurrentPassword string `json:"currpwd" binding:"required"`
	NewPassword     string `json:"newpwd" binding:"required"`
}

func ChangePasswordUser(db *gorm.DB, UpdatedUser UserChangePasswordRequest) (bool, error) {
	var user models.User

	err := db.Find(&user, "uuid = ? AND password = ?", UpdatedUser.UUID, UpdatedUser.CurrentPassword).Error
	isChanged := false
	if user != (models.User{}) {
		fmt.Println(user)
		err = db.Model(&user).Where("uuid = ? AND password = ?", UpdatedUser.UUID, UpdatedUser.CurrentPassword).Update("password", UpdatedUser.NewPassword).Error
		isChanged = true
	}

	return isChanged, err
}
