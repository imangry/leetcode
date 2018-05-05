package main

import (
	"math"
)


func reverse(x int) int {
	tmp := x
	positive := true
	if x < 0 {
		tmp = -x
		positive = false
	}
	var result int
	for tmp != 0 {
		mo := tmp % 10
		result = result*10 + mo
		tmp = tmp / 10

	}

	if !positive {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
