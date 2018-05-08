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

type RandomListNode struct {
	Label  int
	Next   *RandomListNode
	Random *RandomListNode
}

//每个节点拥有一个next，还有一个random 指针
//求深拷贝
//138. Copy List with Random Pointer
func copyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	dummyNode := &RandomListNode{}
	cur := dummyNode
	//旧节点和新节点的对应关系
	randomMap := make(map[*RandomListNode]*RandomListNode)

	for head != nil {
		newNode := &RandomListNode{Label: head.Label}
		cur.Next = newNode
		randomMap[head] = newNode
		newNode.Random = head.Random //新节点的random和旧节点的random一致
		head = head.Next
		cur = cur.Next
	}
	cur = dummyNode.Next
	for cur != nil {
		if cur.Random != nil {
			cur.Random = randomMap[cur.Random]
		}
		cur = cur.Next
	}
	return dummyNode.Next
}

type LRUNode struct {
	Key  int
	Val  int
	Pre  *LRUNode
	Next *LRUNode
}

type LRUCache struct {
	M        map[int]*LRUNode
	Head     *LRUNode
	Tail     *LRUNode
	Capacity int
}

func NewLRUCache(capacity int) LRUCache {
	tail := &LRUNode{}
	head := &LRUNode{
		Pre:  nil,
		Next: tail,
	}
	tail.Pre = head
	return LRUCache{
		M:        make(map[int]*LRUNode),
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {

	if cur, ok := this.M[key]; ok {
		cur.Pre.Next = cur.Next
		cur.Next.Pre = cur.Pre

		this.Tail.Pre.Next = cur
		cur.Pre = this.Tail.Pre
		cur.Next = this.Tail
		this.Tail.Pre = cur
		return cur.Val
	}
	return -1
}

func (this *LRUCache) Set(key, val int) {
	if this.Get(key) != -1 {
		this.M[key].Val = val
		return
	}

	if len(this.M) == this.Capacity {
		delete(this.M, this.Head.Next.Key)
		this.Head.Next = this.Head.Next.Next
		this.Head.Next.Pre = this.Head
	}
	newNode := &LRUNode{Key: key, Val: val}
	this.M[key] = newNode

	this.Tail.Pre.Next = newNode
	newNode.Pre = this.Tail.Pre
	newNode.Next = this.Tail
	this.Tail.Pre = newNode
}
