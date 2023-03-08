package configs

import (
	"go-shop-api/models"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ===============================================
//
// 		Q U E R Y      T O     O R D E R
//
// ===============================================

func GetAllOrder(db *gorm.DB, userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order

	err := db.Find(&orders, "user_id = ?", userID).Error

	return orders, err
}

func GetOrderByID(db *gorm.DB, userID uuid.UUID, orderID uuid.UUID) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	var order models.Order

	err := db.Find(&order, "user_id = ? AND id = ?", userID, orderID).Error

	if order != (models.Order{}) {
		err = db.Find(&orderItems, "order_id = ?", orderID).Error
	}
	return orderItems, err
}

func CancelOrderByID(db *gorm.DB, order models.Order) error {

	err := db.Find(&order, "user_id = ? AND id = ?", order.UserID, order.ID).Update("status", "canceled").Error

	return err
}

func DeleteAllOrderItem(db *gorm.DB, order models.Order) error {
	err := db.Find(&order, "user_id = ? AND id = ?", order.UserID, order.ID).Error

	return err
	// need to be fixed
}

func DeleteOrderByID(db *gorm.DB, order models.Order) error {

	err := db.Where("user_id = ? AND id = ?", order.UserID, order.ID).Delete(&order).Error

	return err
}
