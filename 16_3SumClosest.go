package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println(threeSumClosest([]int{0, 2, 1,-3}, 0))
}

func threeSumClosest(nums []int, target int) (result int) {

	if len(nums) <= 3 {
		for _, value := range nums {
			result += value
		}
		return result
	}

	sort.Ints(nums)

	result = nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; {
		j, k := i+1, len(nums)-1;
		for j < k {

			sum := nums[i] + nums[j] + nums[k]

			if sum == target {
				return sum
			}

			if abs(target-sum) < abs(target-result) {
				result = sum
			}

			if sum > target && j < k {
				k--
			}

			if sum < target && j < k {
				j++
			}
		}
		i++
	}

	return result
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
