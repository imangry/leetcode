package dp

/*
最长非降子序列的长度
eg. 5,3,4,8,6,7,2
i=1 1    1
i=2 1    1
i=3 2    d(2)+1
i=4 3    d(3)+1
i=5 3    d(3)+1
i=6 4    d(4)+1
i=7 1    1

d(i) = max(1,d[j]+1) j<i 且 A[j] <= A[i]
*/

func LongestIncreasingSubsequence(list []int) int {
	if len(list) < 2 {
		return len(list)
	}
	dp := []int{}
	dp = append(dp, 1)

	for i := 1; i < len(list); i++ {
		//默认是1
		max := 1
		for j := 0; j < i; j++ {
			if list[i] >= list[j] && max < dp[j]+1 {
				max = dp[j] + 1
			}
		}
		dp = append(dp, max)
	}
	return dp[len(list)-1]
}
