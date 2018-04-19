package basic_algorithm

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	quickSort(l)
	fmt.Println(l)

	l1 := []int{6, 5, 4, 3, 2, 1}
	quickSort(l1)
	fmt.Println(l1)

	l2 := []int{2, 2, 2, 2, 2, 2}
	quickSort(l2)
	fmt.Println(l2)

	l3 := []int{}
	quickSort(l3)
	fmt.Println(l3)

	l4 := []int{1}
	quickSort(l4)
	fmt.Println(l4)

	l5 := []int{2, 5, 4, 2, 7, -1}
	quickSort(l5)
	fmt.Println(l5)

}
