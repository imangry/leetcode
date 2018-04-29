package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfile := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfile {
			maxProfile = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfile
}
