package dp

//递推公式：
//count[i] = sum(count[j]*count[i-j-1]) 0<=j<=i-1
//对于一个有i节点的BST来说，左边有0...i-1个节点，右边有i-1...0个节点
func TreeNum(n int) int {
	if n < 1 {
		return 0
	}

	count := make([]int, n+1)
	count[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			count[i] += count[j] * count[i-j-1]
		}
	}
	return count[n]
}
