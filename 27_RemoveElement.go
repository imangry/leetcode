package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if nums[j] == val {
			continue
		} else {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
