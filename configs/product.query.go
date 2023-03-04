package configs

import (
	"go-shop-api/models"

	"gorm.io/gorm"
)

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
