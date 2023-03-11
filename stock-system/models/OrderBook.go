package models

// Define a OrderBook struct to list the buyOrder and sellOrder
type OrderBook struct {
	BuyOrder  []BuyOrder
	SellOrder []SellOrder
}
