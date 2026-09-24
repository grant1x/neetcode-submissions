func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]

	for _, p := range prices {
		buy = min(buy, p)
		profit = max(profit, p-buy)
	}

   return profit
}