package main

import "fmt"
type ListNode struct {
	Val  int
	Next *ListNode
}
func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	lists := mergeTwoLists(l1, l2)
	for lists!=nil  {
		fmt.Println(lists.Val)
		lists = lists.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp := new(ListNode)
	result := tmp
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		tmp.Next = l2
	}
	if l2 == nil {
		tmp.Next = l1
	}
	return result.Next
}
