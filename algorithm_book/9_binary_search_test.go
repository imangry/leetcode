package algorithm_book

import (
	"testing"
	"fmt"
)

func TestLowerBound(t *testing.T) {
	list := []int{2, 4, 6, 7}
	fmt.Println(LowerBound(list,8))

}
