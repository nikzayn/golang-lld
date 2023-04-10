package models

type Restaurant struct {
	CustomerID   string
	RestaurantID string
	MenuItems    []Menu
	TotalAmount  int
	Status       int
}
