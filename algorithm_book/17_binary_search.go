package algorithm_book

//求target在数组中第一次出现的位置
//求target的lowerBound
func FirstPositionOfTarget(list []int, target int) int {
	if len(list) < 1 {
		return -1
	}

	left := -1
	right := len(list)

	for left+1 < right {
		mid := (left + right) / 2
		if list[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if right == len(list) || list[right] != target {
		return -1
	} else {
		return right
	}
}
