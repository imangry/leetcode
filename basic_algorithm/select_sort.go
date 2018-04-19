package basic_algorithm
/**
思路：
将当前值视为最小值，然后和后面的值比较，找到最小值，和当前值交换
 */
func SelectSort(list []int) {

	if len(list) < 2 {
		return
	}

	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}

}
