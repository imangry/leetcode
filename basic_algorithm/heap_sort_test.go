package basic_algorithm

/**
堆排序：
nlogn
不稳定
 */
import (
	"testing"
	"fmt"
)

func TestPriorityQueue(t *testing.T) {
	l := []int{5, 6, 3, 8, 1, 4, 7}
	PriorityQueue(l)

	for tmp:=l;; {
		var one int
		tmp, one = ExtractFirst(tmp)
		fmt.Println(tmp, one)
		if len(tmp) < 1 {
			break
		}
	}
}


func TestHeapSort(t *testing.T) {
	l := []int{5, 6, 3, 8, 1, 4, 7}
	HeapSort(l)
	fmt.Println(l)

	l = []int{1, 2, 3, 4, 5, 6}
	HeapSort(l)
	fmt.Println(l)

	l1 := []int{6, 5, 4, 3, 2, 1}
	HeapSort(l1)
	fmt.Println(l1)

	l2 := []int{2, 2, 2, 2, 2, 2}
	HeapSort(l2)
	fmt.Println(l2)

	l3 := []int{}
	HeapSort(l3)
	fmt.Println(l3)

	l4 := []int{1}
	HeapSort(l4)
	fmt.Println(l4)

	l5 := []int{2, 5, 4, 2, 7, -1}
	HeapSort(l5)
	fmt.Println(l5)

}
