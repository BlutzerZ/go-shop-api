package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Email      string    `gorm:"unique;not null"`
	Username   string    `gorm:"unique;not null"`
	Password   string
	DateCreate int64
	DateUpdate int64
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UUID", uuid.NewV4())
	return nil
}
