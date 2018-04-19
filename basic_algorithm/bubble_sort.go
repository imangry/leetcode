package basic_algorithm
/**
思路：
第一次把最大值浮上去，
第二次把第二大的值浮上去
......
第n-1次把第n-1大的值浮上去
 */
func BubbleSort(list []int) {
	if len(list) < 2 {
		return
	}
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
