package main

func maxProfit(prices []int) int {
	profit, minn := 0, 100_000
	for _, p := range prices {
		minn = min(minn, p)
		profit = max(profit, p-minn)
	}
	return profit
}
