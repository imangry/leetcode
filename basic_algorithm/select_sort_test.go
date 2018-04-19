package basic_algorithm

/**
选择排序：
时间复杂度 n^2
不稳定
 */
import (
	"testing"
	"fmt"
)

func TestSelectSort(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	SelectSort(l)
	fmt.Println(l)

	l1 := []int{6, 5, 4, 3, 2, 1}
	SelectSort(l1)
	fmt.Println(l1)

	l2 := []int{2, 2, 2, 2, 2, 2}
	SelectSort(l2)
	fmt.Println(l2)

	l3 := []int{}
	SelectSort(l3)
	fmt.Println(l3)

	l4 := []int{1}
	SelectSort(l4)
	fmt.Println(l4)

	l5 := []int{2, 5, 4, 2, 7, -1}
	SelectSort(l5)
	fmt.Println(l5)

}
