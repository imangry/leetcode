package main

import "fmt"

func main() {
	s := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(s, 0))
}
func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[i] <= nums[mid] {
			if target >= nums[i] && target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}

		if nums[j] >= nums[mid] {
			if target > nums[mid] && target <= nums[j] {
				i = mid + 1
			}else {
				j = mid -1
			}
		}
	}
	return -1
}
