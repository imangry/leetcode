package basic_algorithm

func BinarySearch(list []int, val int) int {

	if len(list) == 0 {
		return -1
	}
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == val {
			return mid
		} else if list[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
