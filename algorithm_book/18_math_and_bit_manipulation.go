package algorithm_book

//判断一个数是否为2的倍数
//i&(i-1)
func CheckPowerOf2(i int) bool {
	if i < 0 {
		return false
	}
	return i&(i-1) == 0
}

//异或：相同则为0，不同为1
//比较两个数不同的比特位个数
//a_xor_b & (a_xor_b - 1) 统计1的个数
func FlipBits(a int, b int) int {
	a_xor_b := a ^ b
	result := 0
	for a_xor_b != 0 {
		a_xor_b = a_xor_b & (a_xor_b - 1)
		result ++
	}
	return result
}

//计算 n! 的结果有多少个后缀0
//思路：计算有多少个5
func TrailingZeroes(n int) int {

	if n < 0 {
		return -1
	}
	count := 0

	for n > 0 {
		n = n / 5
		count += n
	}
	return count
}
