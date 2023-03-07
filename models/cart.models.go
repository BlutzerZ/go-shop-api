package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID     uuid.UUID `gorm:"type:char(36)"`
	DateUpdate int64
}

type CartItem struct {
	CartID     uuid.UUID `gorm:"type:char(36)"`
	Cart       Cart      `gorm:"foreignKey:CartID"`
	ProductID  uuid.UUID `gorm:"type:char(36);primary_key"`
	Qty        int
	DateUpdate int64
}

func (c *Cart) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4())
	return nil
}
