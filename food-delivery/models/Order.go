package models

type Order struct {
	Id       string
	FoodItem FoodItem
	User     User
}
