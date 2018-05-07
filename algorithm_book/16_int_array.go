package algorithm_book

/*
数组的问题：
1.从后向前遍历
2、快慢两个指针
*/

func RemoveElement(list []int, target int) int {

	left, right := 0, 0

	for right < len(list) {

		if list[right] == target {
			right++
		} else {
			list[left] = list[right]
			left ++
			right++
		}
	}
	return left
}

//Remove Duplicates from Sorted Array
//快慢指针
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i, j := 1, 1
	for j < len(nums) {
		if nums[j] == nums[j-1] {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}

func RemoveDuplicates_2(nums []int) int {

	if len(nums) < 3 {
		return len(nums)
	}

	left, right := 1, 2

	for right < len(nums) {
		if nums[right] != nums[left] || nums[right] != nums[left-1] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1

}

//求子数组的和；从0开始到索引i的和为f(i)
//f(i2)-f(i1)=0 => f(i2)=f(i1)
func ZeroSumSubArray(list []int) (int, int) {
	m := make(map[int]int)
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
		if sum == 0 {
			return 0, i
		}

		if pre, ok := m[sum]; ok {
			return pre + 1, i
		} else {
			m[sum] = i
		}
	}
	return 0, 0
}

func PartitionArray(list []int, k int) int {
	if len(list) < 1 {
		return -1
	}

	left := 0
	right := len(list) - 1

	for left < right {
		for left < right && list[right] >= k {
			right --
		}
		for left < right && list[left] < k {
			left ++
		}
		if left < right {
			list[left], list[right] = list[right], list[left]
			left ++
			right --
		}
	}
	return left
}

//TODO
func threeSumClosest(list []int, target int) int {
	return -1
}

func Median(list []int) int {
	mid := (len(list) - 1) / 2
	left := 0
	right := len(list) - 1
	head := list[0]

	for {
		tleft := left
		tright := right
		for left < right {
			for left < right && list[right] > head {
				right --
			}
			if left < right {
				list[left] = list[right]
				left ++
			}

			for left < right && list[left] <= head {
				left ++
			}
			if left < right {
				list[right] = list[left]
				right --
			}
		}
		list[left] = head
		if left == mid {
			return list[left]
		} else if left < mid {
			left++
			head = list[left]
			right = tright
		} else {
			head = list[tleft]
			right = left-1
			left = tleft
		}
	}
}
