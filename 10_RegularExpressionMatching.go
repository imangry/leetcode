package main

func main() {

}
func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}
	var dp [len(s) + 1][len(p) + 1]bool
	dp[0][0] = true

	for j := 0; j < len(p); j++ {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}
	return false
}
