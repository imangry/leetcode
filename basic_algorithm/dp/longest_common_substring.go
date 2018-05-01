package dp

func LongestCommonSubstring(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	dp := make([][]int, len(s1)+1)

	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	max := 0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			}
		}
	}
	return max
}
