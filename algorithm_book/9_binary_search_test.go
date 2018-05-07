package algorithm_book

import (
	"testing"
	"fmt"
)

func TestLowerBound(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(LowerBound(list, 3))
}
