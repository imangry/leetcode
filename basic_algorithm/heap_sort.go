package basic_algorithm

/*
大顶堆
使得所有的子节点都小于父节点
因此从 n/2-1这个最小的父节点开始进行调整
parent = (son-1)/2
left = parent*2+1
right = parent*2+2
*/
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
	k := len(list)
	list = append(list, i)

	SiftUp(list, k)
	return list
}

func SiftUp(list []int, k int) {
	i := list[k]
	for k > 0 {
		parent := (k - 1) / 2
		if list[parent] > list[k] {
			break
		}
		list[k] = list[parent]
		k = parent
	}
	list[k] = i
}

func SiftDown(list []int, i int) {
	tmp := list[i]
	for j := 2*i + 1; j < len(list); j = 2*j + 1 {
		if j+1 < len(list) && list[j+1] > list[j] {
			j++
		}
		if list[j] > tmp {
			list[i] = list[j]
			i = j
		} else {
			break
		}
	}
	list[i] = tmp
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
