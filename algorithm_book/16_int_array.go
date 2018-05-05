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
		for left < right && list[right] > k {
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
