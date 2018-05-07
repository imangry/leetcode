package algorithm_book

import (
	"math"
	"fmt"
)

/*
1.前提是**有序的数组**
2.跳出循环条件：left right相遇；mid命中了
*/

//大于等于目标值的最小索引
func LowerBound(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}
	left := -1
	right := len(list)

	for left+1 < right {
		mid := (left + right) / 2

		if list[mid] >= target {
			right = mid
		} else {
			left = mid
		}
		fmt.Println("left", left, "right", right)
	}
	return left + 1
}

func UpperBound(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}
	left := -1
	right := len(list)

	for left+1 < right {
		mid := (left + right) / 2
		if list[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	return right - 1
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
