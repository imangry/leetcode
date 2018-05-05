package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; {
		sum := -nums[i]

		for j, k := i+1, len(nums)-1; j < k; {

			if nums[j]+nums[k] < sum {
				j++
			}
			if nums[j]+nums[k] > sum {
				k--
			}
			if nums[j]+nums[k] == sum {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				for nums[k] == nums[k+1] && j < k {
					k--
				}
			}

		}
		i++
		for nums[i] == nums[i-1] && i < len(nums)-2 {
			i++
		}
	}
	return result
}
