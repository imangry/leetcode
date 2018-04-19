package basic_algorithm

/**
思路：
从右到左找到第一个比head小的，交换
从左到右找到第一个比head大的，交换

注意left<right的条件
 */
func QuickSort(list []int) {
	if len(list) < 2 {
		return
	}
	head := list[0]
	left := 0
	right := len(list) - 1

	for left < right {
		for left < right && list[right] > head {
			right--
		}
		if left < right {
			list[left] = list[right]
			left++
		}
		for left < right && list[left] < head {
			left++
		}
		if left < right {
			list[right] = list[left]
			right--
		}
	}
	list[left] = head

	QuickSort(list[0:left])
	QuickSort(list[left+1:])
}
