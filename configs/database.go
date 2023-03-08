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

	db.AutoMigrate(&models.Cart{})
	db.AutoMigrate(&models.CartItem{})

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})

	return db
}

var DB *gorm.DB = ConnectDB()
