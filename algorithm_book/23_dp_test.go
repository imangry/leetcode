package algorithm_book

import (
	"testing"
	"fmt"
)

func TestTriangle(t *testing.T) {
	in := make([][]int, 4)
	in[0] = []int{2}
	in[1] = []int{3, 4}
	in[2] = []int{6, 5, 7}
	in[3] = []int{4, 1, 8, 3}

	fmt.Println(Triangle(in))

}
