package models

// Define Order struct with detials related to stock, quantity, price etc.
type Order struct {
	Id        string
	Timestamp string
	Name      string
	Action    string
	Price     float64
	Quantity  int
}
