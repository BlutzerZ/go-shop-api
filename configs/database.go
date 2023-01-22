package configs

import (
	"go-shop-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/goshopapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.ProductCategory{})

	return db
}

var DB *gorm.DB = ConnectDB()

func AddUser(db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func DeleteUser(db *gorm.DB, ID int) (int, error) {
	err := db.Delete(models.User{}, ID).Error

	return ID, err
}
