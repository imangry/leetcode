package algorithm_book

import (
	"fmt"
	"testing"
)

func TestRotateString(t *testing.T) {
	fmt.Println(RotateString("abcde", 2))
}

func TestReverse(t *testing.T) {
	bytes := []byte{100,101,102}
	Reverse(bytes)
	fmt.Println(bytes)
}
