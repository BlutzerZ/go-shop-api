package configs

import (
	"fmt"
	"go-shop-api/models"
	"log"
	"os"

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

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.ProductCategory{})

	return db
}

var DB *gorm.DB = ConnectDB()

// ==================================================
//
//              S O M E    Q U E R Y
//
// ==================================================

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

func DeleteUser(db *gorm.DB, ID int) (int, error) {
	err := db.Delete(models.User{}, ID).Error

	return ID, err
}

// ====================
//   PRODUCT QUERY
// ====================

func GetProductByLimit(db *gorm.DB, limit int) ([]models.Product, error) {
	var products []models.Product

	err := db.Limit(limit).Find(&products).Error

	return products, err
}

func GetProductByID(db *gorm.DB, ID int) (models.Product, error) {
	var product models.Product

	err := db.Find(&product, ID).Error

	return product, err
}

func AddProduct(db *gorm.DB, product models.Product) (models.Product, error) {
	err := db.Create(&product).Error

	return product, err
}
