package algorithm_book

//奇偶交换两个node
func SwapNodesInPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next.Next
	next := head.Next
	next.Next = head
	head.Next = SwapNodesInPairs(tmp)
	return next
}
