package basic_algorithm

func quickSort(list []int) {
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

	quickSort(list[0:left])
	quickSort(list[left+1:])
}
