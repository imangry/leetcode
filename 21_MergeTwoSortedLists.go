package main

type ListNode struct {
	Val  int
	Next *ListNode
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
