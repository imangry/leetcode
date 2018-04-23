package basic_algorithm

import (
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {

	l5 := []int{2, 5, 4, 2, 7, -1}
	BubbleSort(l5)
	fmt.Println(l5)
	fmt.Println(BinarySearch(l5, 8))

}
