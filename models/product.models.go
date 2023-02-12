package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	Name       string
	Desc       string
	Stock      int
	CatID      int
	DateCreate int64
	DateUpdate int64
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4())
	return nil
}

type ProductCategory struct {
	CatID int
	Name  string
}
