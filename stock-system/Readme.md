# Stock System
- In this repository, we are going to implement a stock exchange system where we need to replicate the real life
example of stock exchange system like to buy and sell the stocks

## Overview
- Design Thoughts
- Usage
- Technologies Required
- API Usage
- Demo

### Problem statement
- Traders place Buy or Sell orders for a stock indicating the quantity and price.
- These orders get entered into the exchange’s order-book and remain there until they are matched or until the trading day ends.
- The exchange follows a FirstInFirstOut Price-Time order-matching rule, which states that: “The first order in the order-book at  a price level is the first order matched. All orders at the same price level are filled according to time priority”.
- The exchange works like a market, lower selling prices and higher buying prices get priority.
- A trade is executed when a buy price is greater than or equal to a sell price. The trade is recorded at the price of the sell order regardless of the price of the buy order.

### Approach
- First, I will create a list of structs to modularize my data like `Order`, `BuyOrder`, `SellOrder` and `OrderBook`
- In `Order` struct, I will store fields of data like `Id`, `Timestamp`, `Side`, `Quantity`, `Price`, `Symobol`, `Action`.
- In `BuyOrder` struct, it will store the order related to BUY.
- In `SellOrder` struct, it will store the order related to SELL.
- in `OrderBook` struct, it stores both `BuyOrder` and `SellOrder` struct.
- After all the this, I will create one function to add order whether it's action of `SELL` or `BUY`.
- Create a function `matchOrder` which would check if BUY order is more than equal to SELL order, if it is then execute the
trade and pop that trade, else traverse over other orders


