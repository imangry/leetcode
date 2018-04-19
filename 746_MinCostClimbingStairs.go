package main

func main() {

}
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1],dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1],dp[len(cost)-2])
}

func min(i, j int) int {
	if i < j {
		return i
	}else {
		return j
	}
}