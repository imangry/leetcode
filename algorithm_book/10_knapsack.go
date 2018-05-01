package algorithm_book

import "fmt"

/*
01背包
i代表第几个物品，value代表价值
公式：f(i,weight) = max{ f(i-1,weight), f(i-1,weight-w[i])+value[i] }
时间复杂度：O(N*V)
空间复杂度：O(N*V)
*/
func KnapsackWithoutRepetition(maxWeight int, weight []int, value []int) int {
	dp := make([][]int, len(weight))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, maxWeight)
	}
	//是否放第一个物品；第一个物品的情况下，不同maxWeight的最大价值
	for i := 0; i < maxWeight; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j < maxWeight; j++ {
			//不放当前物品
			dp[i][j] = dp[i-1][j]
			//如果当前容量大于当前物品，看放进去该物品是不是更大
			if j >= weight[i] && (dp[i-1][j-weight[i]]+value[i]) > dp[i][j] {
				dp[i][j] = dp[i-1][j-weight[i]] + value[i]
			}
		}
	}
	for _, value := range dp {
		fmt.Println(value)
	}
	return dp[len(weight)-1][maxWeight-1] + 1
}

/*
01背包使用滚动数组
空间复杂度变为：O(2*N)
 */
func KnapsackWithoutRepetition_1(maxWeight int, weight []int, value []int) int {

	dp := [2][]int{}

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, maxWeight)
	}
	//是否放第一个物品；第一个物品的情况下，不同maxWeight的最大价值
	for i := 0; i < maxWeight; i++ {
		if i >= weight[0] {
			dp[0][i] = value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for j := 0; j < maxWeight; j++ {
			k := i & 1
			//不放当前物品
			dp[k][j] = dp[k^1][j]
			//如果当前容量大于当前物品，看放进去该物品是不是更大
			if j >= weight[i] && (dp[k^1][j-weight[i]]+value[i]) > dp[k][j] {
				dp[k][j] = dp[k^1][j-weight[i]] + value[i]
			}
		}
	}
	max := dp[0][maxWeight-1]
	if dp[1][maxWeight-1] > max {
		max = dp[1][maxWeight-1]
	}
	return max + 1
}
