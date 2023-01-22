package models

type Product struct {
	ID         int
	Name       string
	Desc       string
	Stock      int
	CatID      int
	DateCreate int
	DateUpdate int
}

type ProductCategory struct {
	CatID int
	Name  string
}
