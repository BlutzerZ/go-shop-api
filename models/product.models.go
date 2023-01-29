package models

type Product struct {
	ID         int
	Name       string
	Desc       string
	Stock      int
	CatID      int
	DateCreate int64
	DateUpdate int64
}

type ProductCategory struct {
	CatID int
	Name  string
}
