package main

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); {
		if nums[i] == nums[j] {
			j++
			continue
		}
		i++
		nums[i] = nums[j]
		j++
	}
	return i + 1
}
