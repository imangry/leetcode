package basic_algorithm

func HeapSort(list []int) {
	if len(list) < 2 {
		return
	}

	for i := len(list)/2 - 1; i >= 0; i-- {
		maxHeapify(list, i)
	}
	for i := len(list) - 1; i > 0; i-- {
		list[i], list[0] = list[0], list[i]
		maxHeapify(list[:i], 0)
	}
}

func PriorityQueue(list []int) {
	if len(list) < 2 {
		return
	}
	for i := len(list)/2 - 1; i >= 0; i-- {
		maxHeapify(list, i)
	}
}
func ExtractFirst(list []int) ([]int, int) {
	if len(list) == 0 {
		return list, 0
	}
	if len(list) == 1 {
		return []int{}, list[0]
	}
	tmp := list[0]

	list[0] = list[len(list)-1]
	maxHeapify(list[:len(list)-1], 0)
	return list[:len(list)-1], tmp
}
func Insert(list []int, i int) []int {
	return nil
}

func maxHeapify(list []int, i int) {
	tmp := list[i]

	for k := i*2 + 1; k < len(list); k = k*2 + 1 {
		if k+1 < len(list) && list[k] < list[k+1] {
			k++
		}
		if list[k] > tmp {
			list[i] = list[k]
			i = k
		} else {
			break
		}
	}
	list[i] = tmp
}
