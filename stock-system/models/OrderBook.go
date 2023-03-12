package models

import (
	"fmt"
	"sort"
)

// Define a OrderBook struct to list the buyOrder and sellOrder
type OrderBook struct {
	BuyOrder  []BuyOrder
	SellOrder []SellOrder
}

const BUY_ACTION = "BUY"
const SELL_ACTION = "SELL"

/**
 * Description: In this function, we are going to add order accordingly to the action type
 * @Input - order struct
 * TC - O(1) | SC - O(1)
 */
func (ob *OrderBook) AddOrder(order Order) {
	//Append buy order if action is BUY
	if order.Action == BUY_ACTION {
		ob.BuyOrder = append(ob.BuyOrder, BuyOrder{order})
		//Basically, we need to sort the BuyOrder struct to get the price in high to low order
		sort.Slice(ob.BuyOrder, func(i, j int) bool {
			return ob.BuyOrder[i].Price > ob.BuyOrder[j].Price
		})
	} else if order.Action == SELL_ACTION {
		ob.SellOrder = append(ob.SellOrder, SellOrder{order})
		//Basically, we need to sort the SellOrder struct to get the price low to high order
		sort.Slice(ob.SellOrder, func(i, j int) bool {
			return ob.SellOrder[i].Price < ob.SellOrder[j].Price
		})
	} else {
		return
	}
}

/**
 * Description: In this function, we are going to add order accordingly to the action type
 * @Input - order struct
 * TC - O(1) | SC - O(1)
 */
func (ob *OrderBook) MatchOrders() {
	//Traverse ovver length of buyOrder and sellOrders until it's greater than zero
	for len(ob.BuyOrder) > 0 && len(ob.SellOrder) > 0 {
		if ob.BuyOrder[0].Price < ob.SellOrder[0].Price {
			break
		}
		if ob.BuyOrder[0].Quantity <= ob.SellOrder[0].Quantity {
			fmt.Printf("%s %d %.2f %s\n", ob.SellOrder[0].Id, ob.BuyOrder[0].Quantity, ob.SellOrder[0].Price, ob.BuyOrder[0].Id)
			ob.SellOrder[0].Quantity -= ob.BuyOrder[0].Quantity
			ob.BuyOrder = ob.BuyOrder[1:]
			if ob.SellOrder[0].Quantity == 0 {
				ob.SellOrder = ob.SellOrder[1:]
			}
		} else {
			fmt.Printf("%s %d %.2f %s\n", ob.SellOrder[0].Id, ob.SellOrder[0].Quantity, ob.SellOrder[0].Price, ob.BuyOrder[0].Id)
			ob.BuyOrder[0].Quantity -= ob.SellOrder[0].Quantity
			ob.SellOrder = ob.SellOrder[1:]
		}
	}
}
