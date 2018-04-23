package basic_algorithm

type Node struct {
	Val  int
	Next *Node
}

func reverseListByIteration(head *Node) *Node {

	var pre *Node = nil

	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//TODO 递归反转链表
func reverseListByRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newHead := reverseListByRecursive(next)
	next.Next = head
	head.Next = nil
	return newHead
}

func hasCircle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && slow != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
		if slow == nil {
			break
		}
		if fast == slow {
			return true
		}
	}
	return false
}
