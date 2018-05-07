package algorithm_book

//判断一个数是否为2的倍数
//i&(i-1)
func CheckPowerOf2(i int) bool {
	if i < 0 {
		return false
	}
	return i&(i-1) == 0
}

func FlipBits(a int, b int) int {
	a_xor_b := a ^ b
	result := 0
	for a_xor_b != 0 {
		a_xor_b = a_xor_b & (a_xor_b - 1)
		result ++
	}
	return result
}

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
