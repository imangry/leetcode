package main

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != -1 {
			if nums[nums[i]-1] != -1 {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], -1
				i--
			}
		}
	}

	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != -1 {
			result = append(result, i+1)
		}
	}
	return result
}
