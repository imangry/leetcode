package algorithm_book

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicatesFromSortedList(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}
	cur := root
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return root
}

func RemoveDuplicatesFromSortedList_2(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	dummyNode := &ListNode{Next: root}
	pre := dummyNode
	for root != nil && root.Next != nil {
		if root.Val != root.Next.Val {
			pre = root
			root = root.Next
		} else {
			for root.Next != nil && root.Val == root.Next.Val {
				root.Next = root.Next.Next
			}
			pre.Next = root.Next
			root = pre.Next
		}

	}
	return dummyNode.Next
}
//将链表从target中间分开
func PartitionList(head *ListNode, target int) *ListNode {
	if head == nil {
		return nil
	}
	leftDummy := &ListNode{}
	l := leftDummy
	rightDummy := &ListNode{}
	r := rightDummy

	for head != nil {
		if head.Val < target {
			l.Next = head
			l = l.Next
		} else {
			r.Next = head
			r = r.Next
		}
		head = head.Next
	}
	r.Next = nil
	l.Next = rightDummy.Next
	return leftDummy.Next
}
