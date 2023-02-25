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

type UserChangePasswordRequest struct {
	ID              int    `json:"id" binding:"required"`
	CurrentPassword string `json:"currpwd" binding:"required"`
	NewPassword     string `json:"newpwd" binding:"required"`
}

func ChangePasswordUser(db *gorm.DB, UpdatedUser UserChangePasswordRequest) (bool, error) {
	var user models.User

	err := db.Find(&user, "id = ? AND password = ?", UpdatedUser.ID, UpdatedUser.CurrentPassword).Error
	isChanged := false
	if user != (models.User{}) {
		fmt.Println(user)
		err = db.Model(&user).Where("id = ? AND password = ?", UpdatedUser.ID, UpdatedUser.CurrentPassword).Update("password", UpdatedUser.NewPassword).Error
		isChanged = true
	}

	return isChanged, err
}

// ====================
//   PRODUCT QUERY
// ====================

func GetProductByLimit(db *gorm.DB, limit int) ([]models.Product, error) {
	var products []models.Product

	err := db.Limit(limit).Find(&products).Error

	return products, err
}

func GetProductByID(db *gorm.DB, ID string) (models.Product, error) {
	var product models.Product

	err := db.Find(&product, "id = ?", ID).Error

	return product, err
}

func AddProduct(db *gorm.DB, product models.Product) (models.Product, error) {
	err := db.Create(&product).Error

	return product, err
}

func DeleteProduct(db *gorm.DB, ID string) error {

	err := db.Delete(models.Product{}, "id = ?", ID).Error

	return err
}

func UpdateProduct(db *gorm.DB, ID string, updateProduct models.Product) (models.Product, error) {
	var product models.Product

	err := db.Model(&product).Where("id = ?", ID).Updates(updateProduct).Error

	return updateProduct, err

}

// ====================
//   CART QUERY
// ====================

func InsertItemToCart(db *gorm.DB, userID int, item models.CartItem) error {
	var cart models.Cart

	// First find user_id in cart model
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	if cart != (models.Cart{}) {
		// then use cart.ID to fill item.CartID
		item.CartID = cart.ID
		// and finally add it on CartItem models
		err = db.Create(&item).Error

		// this if first time to add item on cart
	} else {
		cart.UserID = userID
		err = db.Create(&cart).Error
		if err != nil {
			return err
		}
		item.CartID = cart.ID
		err = db.Create(&item).Error
	}

	return err
}

func EditItemCart(db *gorm.DB, userID int, item models.CartItem) error {
	var cart models.Cart
	var cartItem models.CartItem

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}
	// Update item on cart
	err = db.Model(&cartItem).Where("cart_id = ? and product_id =?", cart.ID, item.ProductID).Updates(&item).Error

	return err
}

func DeleteItemcART(db *gorm.DB, userID int, item models.CartItem) error {
	var cart models.Cart

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	// Delete item ont cart
	err = db.Where("cart_id = ?", cart.ID).Delete(&item).Error

	return err
}

func GetAllItemCart(db *gorm.DB, userID int) error {
	var cart models.Cart
	var cartItems []models.CartItem

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	// Show all item on cart
	err = db.Find(&cartItems).Where("cart_id = ?", cart.ID).Error

	return err
}

func ForwardCartToOrder(db *gorm.DB, userID int) error {
	var cart models.Cart

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	// forward
	// err = db.Create(&)
}
