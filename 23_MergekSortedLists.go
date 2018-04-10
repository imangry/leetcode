package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{5, 3, 2, 45, 89}))
}


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	l1 := mergeKLists(lists[0:mid])
	l2 := mergeKLists(lists[mid:])
	return merge(l1, l2)
}
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}

func mergeSort(list []int) []int {
	if len(list) == 1 || len(list) == 0 {
		return list
	}

	mid := len(list) / 2
	mergeSort(list[0:mid])
	mergeSort(list[mid:])

	return merge1(list[0:mid], list[mid:])

}
func merge1(left, right []int) []int {
	tmp := make([]int, 0)

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			tmp = append(tmp, left[0])
			left = left[1:]
		} else {
			tmp = append(tmp, right[0])
			right = right[1:]
		}
	}
	tmp = append(tmp, left...)
	tmp = append(tmp, right...)

	return tmp
}
