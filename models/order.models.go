package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:char(12);primary_key"`
	UserID     uuid.UUID `gorm:"type:char(24)"`
	CartID     uuid.UUID `gorm:"type:char(12)"`
	Adress     string
	DateUpdate int64
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4())
	return nil
}
