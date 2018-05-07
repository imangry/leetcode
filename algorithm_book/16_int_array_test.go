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

func TestPartitionArray(t *testing.T) {
	fmt.Println(PartitionArray([]int{3, 2, 2, 1}, 2))
}

func TestMedian(t *testing.T) {
	fmt.Println(Median([]int{3,4,1,2,5,6}))
}
