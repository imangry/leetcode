package algorithm_book

import "fmt"

/*
from left to right
*/
func RotateString(s string, k int) string {

	a := []byte(s)

	k = k % len(s)
	fmt.Println(a)

	Reverse(a[:len(a)-k])
	Reverse(a[len(a)-k:])
	Reverse(a)

	return string(a)
}
func Reverse(list []byte) {
	l := len(list)
	mid := l / 2
	for i := 0; i < mid; i++ {
		list[i], list[l-1-i] = list[l-1-i], list[i]
	}
}
