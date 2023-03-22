package main

import (
	"strings"

	helpers "github.com/nikzayn/lld-golang/stock-system/helpers"
	models "github.com/nikzayn/lld-golang/stock-system/models"
)

func main() {
	orders := []string{
		"#1 09:45 BAC SELL 100 240.10",
		"#2 09:45 BAC SELL 90 237.45",
		"#3 09:47 BAC BUY 80 238.10",
		"#5 09:48 BAC SELL 220 241.50",
		"#6 09:49 BAC BUY 50 238.50",
		"#7 09:52 TCS BUY 10 1001.10",
		"#8 10:01 BAC SELL 20 240.10",
		"#9 10:02 BAC BUY 150 242.70",
	}

	orderBook := models.OrderBook{}

	for _, orderStr := range orders {
		parts := strings.Split(orderStr, " ")
		order := models.Order{
			Id:        parts[0],
			Timestamp: parts[1],
			Name:      parts[2],
			Action:    parts[3],
			Quantity:  helpers.ConvertStringToInt(parts[4]),
			Price:     helpers.ConvertStringToFloat(parts[5]),
		}
		orderBook.AddOrder(order)
		orderBook.MatchOrders()
	}
}
