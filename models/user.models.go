package models

type User struct {
	ID         int
	Email      string
	Username   string
	Password   string
	DateCreate int64
	DateUpdate int64
}
