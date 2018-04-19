package main

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfile := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		maxProfile = max(maxProfile, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfile
}
