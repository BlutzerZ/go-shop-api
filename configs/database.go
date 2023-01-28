package configs

import (
	"go-shop-api/models"
	"log"
	"os"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	godotenv.Load(".env")

	DBUsername := os.Getenv("DBUSERNAME")
	DBPassword := os.Getenv("DBPASSWORD")
	DBHost := os.Getenv("DBHOST")
	DBPort := os.Getenv("DBPORT")
	DBName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBHost, DBPort, DBName)
	// dsn := "blutzerz:unknown7703;@tcp(mysqlblutzerz.mysql.database.azure.com:3306)/goshopapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&models.Product{}, &models.User{}, &models.ProductCategory{})

	return db
}

var DB *gorm.DB = ConnectDB()

// ==================================================
//
//              S O M E    Q U E R Y
//
// ==================================================

func AddUser(db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error

	return user, err
}

func DeleteUser(db *gorm.DB, ID int) (int, error) {
	err := db.Delete(models.User{}, ID).Error

	return ID, err
}
