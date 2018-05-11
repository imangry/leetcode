package algorithm_book

func Triangle(in [][]int) int {
	if len(in) == 0 {
		return 0
	}
	dp := make([]int, len(in[len(in)-1]))
	for i := 0; i < len(dp); i++ {
		dp[i] = in[len(in)-1][i]
	}
	for i := len(in) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if in[i][j]+dp[j] < in[i][j]+dp[j+1] {
				dp[j] = in[i][j] + dp[j]
			} else {
				dp[j] = in[i][j] + dp[j+1]
			}
		}
	}
	return dp[0]
}
