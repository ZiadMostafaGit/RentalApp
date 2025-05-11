package models

type Item struct {
	Id          uint
	Owner_id    uint
	Title       string
	Description string
	Price       int
	Status      string
}

type Item_image struct {
	Item_id    uint
	Item_image string
}
