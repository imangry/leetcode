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
