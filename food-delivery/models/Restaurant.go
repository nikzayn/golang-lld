package models

type Restaurant struct {
	Id            string
	Name          string
	ListOfPincode []string
	FoodItem      FoodItem
	Ratings       Rating
}
