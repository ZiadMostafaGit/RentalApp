package models

type Categories struct {
	Id   uint
	Name string
}

type Item_in_category struct {
	Item_id     uint
	Category_id uint
}
