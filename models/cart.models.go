package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID `gorm:"type:char(12);primary_key"`
	UserID     int       `gorm:"type:char(24)"`
	DateUpdate int64
}

type CartItem struct {
	CartID     uuid.UUID `gorm:"type:char(12);primary_key"`
	ProductID  uuid.UUID `gorm:"type:char(36)"`
	Qty        int
	DateUpdate int64
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4())
	return nil
}