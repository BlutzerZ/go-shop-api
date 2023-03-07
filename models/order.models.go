package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID     uuid.UUID `gorm:"type:char(36)"`
	Adress     string
	Status     string
	TOtal      int
	Discount   int
	DateUpdate int64
}

type OrderItem struct {
	OrderID   uuid.UUID `gorm:"type:char(36)"`
	Order     Order     `gorm:"foreignKey:OrderID"`
	ProductID uuid.UUID `gorm:"type:char(36);primary_key"`
	Qty       int
	Price     int
	Discount  int
	SubTotal  int
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4())
	return nil
}
