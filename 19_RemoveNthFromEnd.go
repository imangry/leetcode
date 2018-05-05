package main


func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		tmp := head.Next
		head.Next = nil
		return tmp
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	tmp := slow.Next.Next
	slow.Next.Next = nil
	slow.Next = tmp
	return head
}
