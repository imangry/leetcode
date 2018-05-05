package main



func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head.Next
	head.Next = swapPairs(first.Next.Next)
	first.Next = head
	return first
}
func swapPairs_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode{Next: head}

	current := tmp

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		first.Next = second.Next
		second.Next = first
		current.Next = second
		current = current.Next.Next.Next
	}
	return tmp.Next
}
