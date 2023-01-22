package models

type User struct {
	ID         int
	Email      string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	Password   string
	DateCreate int64
	DateUpdate int64
}
