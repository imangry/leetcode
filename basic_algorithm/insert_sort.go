package basic_algorithm

func InsertSort(list []int) {
	if len(list) < 2 {
		return
	}
	for i := 1; i < len(list); i++ {
		cur := list[i]
		for j := i - 1; j >= 0; j-- {
			if list[j] > cur {
				list[j], list[j+1] = list[j+1], list[j]
			} else {
				list[j+1] = cur
				break
			}
		}
	}
}
