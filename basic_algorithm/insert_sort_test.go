package basic_algorithm

/**
插入排序：
稳定
最好 n
最坏 n^2
平均 n^2
 */
import (
	"testing"
	"fmt"
)

func TestInsertSort(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	InsertSort(l)
	fmt.Println(l)

	l1 := []int{6, 5, 4, 3, 2, 1}
	InsertSort(l1)
	fmt.Println(l1)

	l2 := []int{2, 2, 2, 2, 2, 2}
	InsertSort(l2)
	fmt.Println(l2)

	l3 := []int{}
	InsertSort(l3)
	fmt.Println(l3)

	l4 := []int{1}
	InsertSort(l4)
	fmt.Println(l4)

	l5 := []int{2, 5, 4, 2, 7, -1}
	InsertSort(l5)
	fmt.Println(l5)

}
