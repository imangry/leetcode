package algorithm_book

import (
	"testing"
	"fmt"
)

func TestRemoveElement(t *testing.T) {

	list := []int{3, 2, 3, 2, 4, 3}
	fmt.Println(list)
	fmt.Println(RemoveElement(list, 3))
}
