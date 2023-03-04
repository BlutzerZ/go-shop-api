package configs

import (
	"go-shop-api/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ====================
//   CART QUERY
// ====================

func InsertItemToCart(db *gorm.DB, UserID uuid.UUID, item models.CartItem) error {
	var cart models.Cart

	// First find user_id in cart model
	err := db.Find(&cart, "user_id = ?", UserID).Error
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
		cart.UserID = UserID
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

	return err
}
