package basic_algorithm

/**
快速排序：
非稳定
平均时间复杂度 nlog(n)
最坏时间复杂度 n^2
 */
import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	QuickSort(l)
	fmt.Println(l)

	l1 := []int{6, 5, 4, 3, 2, 1}
	QuickSort(l1)
	fmt.Println(l1)

	l2 := []int{2, 2, 2, 2, 2, 2}
	QuickSort(l2)
	fmt.Println(l2)

	l3 := []int{}
	QuickSort(l3)
	fmt.Println(l3)

	l4 := []int{1}
	QuickSort(l4)
	fmt.Println(l4)

	l5 := []int{2, 5, 4, 2, 7, -1}
	QuickSort(l5)
	fmt.Println(l5)

}
