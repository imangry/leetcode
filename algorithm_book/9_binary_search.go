package algorithm_book

import "math"

/*
1.前提是**有序的数组**
2.跳出循环条件：left right相遇；mid命中了
*/

//TODO
func LowerBound(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}
	left := -1
	right := len(list)

	for right-left > 1 {

		mid := (left + right) / 2

		if list[mid] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return left + 1
}

func Sqrt(val float64) float64 {

	left := float64(0)
	right := val
	mid := (left + right) / 2
	for math.Abs(right-left) > 10^(-6) {
		mid = (right - left) / 2
		if mid*mid == val {
			return mid
		} else if mid*mid > val {
			right = mid
		} else {
			left = mid
		}
	}
	return mid
}
