package algorithm_book

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	list := []int{3, 2, 3, 2, 4, 3}
	fmt.Println(list)
	fmt.Println(RemoveElement(list, 3))
}

func TestZeroSumSubArray(t *testing.T) {
	fmt.Println(ZeroSumSubArray([]int{2, 3, -3, 1}))
}
