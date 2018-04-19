package main

import (
	"fmt"
)

func main() {
	s := []int{6, 5, 4, 2,3}
	nextPermutation(s)
	fmt.Println(s)
}
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	k := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		reverseInt(nums)
		return
	}

	i := len(nums) - 1
	for ; i >= 0; i-- {
		if nums[k] < nums[i] {
			break
		}
	}
	tmp := nums[k]
	nums[k] = nums[i]
	nums[i] = tmp

	reverseInt(nums[k+1:])
}

func reverseInt(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
}
