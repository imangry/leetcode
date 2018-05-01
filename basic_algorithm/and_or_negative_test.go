package basic_algorithm

import (
	"fmt"
	"testing"
)

func TestAnd(t *testing.T) {

	fmt.Println(1 & 1)
	fmt.Println(1 | 1)
	fmt.Println(1 | 0)
	fmt.Println(1 & 0)
	fmt.Println(^1)
	fmt.Println(^0)

}
