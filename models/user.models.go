package models

type User struct {
	ID         int    `gorm:"primary_key"`
	Email      string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	Password   string
	DateCreate int64
	DateUpdate int64
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	tx.Statement.SetColumn("ID", uuid.NewV4())
// 	return nil
// }
