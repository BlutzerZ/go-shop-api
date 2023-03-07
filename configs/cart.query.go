package configs

import (
	"go-shop-api/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ====================
//   CART QUERY
// ====================

func InsertItemToCart(db *gorm.DB, UserID uuid.UUID, item *models.CartItem) error {
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

func EditItemCart(db *gorm.DB, userID uuid.UUID, item models.CartItem) error {
	var cart models.Cart
	var cartItem models.CartItem

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	// Update item on cart
	err = db.Model(&cartItem).Where("cart_id = ? and product_id = ?", cart.ID, item.ProductID).Updates(&item).Error

	return err
}

func DeleteItemCart(db *gorm.DB, userID uuid.UUID, item models.CartItem) error {
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

func DeleteAllItemCart(db *gorm.DB, userID uuid.UUID) error {
	var cart models.Cart
	var cartItems []models.CartItem

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error
	if err != nil {
		return err
	}

	// Delete item ont cart
	err = db.Where("cart_id = ?", cart.ID).Delete(&cartItems).Error

	return err
}

func GetAllItemCart(db *gorm.DB, userID uuid.UUID) ([]models.CartItem, error) {
	var cart models.Cart
	var cartItems []models.CartItem

	// Get Cart model with user ID
	err := db.Find(&cart, "user_id = ?", userID).Error

	// Show all item on cart
	err = db.Find(&cartItems).Where("cart_id = ?", cart.ID).Error

	return cartItems, err
}

func ForwardCartToOrder(db *gorm.DB, userID uuid.UUID, cartItems []models.CartItem) error {
	var order models.Order
	var orderitems []models.OrderItem

	// create order table
	order.UserID = userID
	order.Status = "payment"
	err := db.Create(&order).Error
	if err != nil {
		return err
	}

	// forward
	for _, cartItem := range cartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Qty:       cartItem.Qty,
			Price:     cartItem.Price,
			Discount:  cartItem.Discount,
			SubTotal:  cartItem.Price * cartItem.Discount,
		}
		orderitems = append(orderitems, orderItem)
	}

	err = db.Create(&orderitems).Error

	return err
}
