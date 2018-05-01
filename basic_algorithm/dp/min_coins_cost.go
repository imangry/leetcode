package dp

import (
	"math"
	"fmt"
)

/*
1,3,5元三种硬币
求组成11元的最少硬币数是多少
每个值都是由之前的一个值加上 1 3 5 得到的
dp[i] = min{ dp[i-coins[j]] + 1 }
*/

func MinCoinsCost(v int) int {
	if v < 1 {
		return 0
	}
	coins := []int{1, 3, 5}
	dp := make([]int, v+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt64
	}
	dp[1] = 1
	dp[3] = 1
	dp[5] = 1
	for i := 2; i <= v; i++ {
		min := dp[i]
		for _, coin := range coins {
			if i-coin > 0 {
				if dp[i-coin]+1 < min {
					min = dp[i-coin] + 1
				}
			}
		}
		dp[i] = min
	}
	fmt.Println(dp)
	return dp[v]
}
