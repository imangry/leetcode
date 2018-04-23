package basic_algorithm

const (
	BLACK = iota
	RED
)

type RBNode struct {
	Val   int
	Left  *RBNode
	Right *RBNode
	RB    int
}

func NewRBNode(val int, left, right *RBNode, rb int) *RBNode {
	return &RBNode{
		Val:   val,
		Left:  left,
		Right: right,
		RB:    rb,
	}
}
func RBInit() *RBNode {
	null := NewRBNode(0, nil, nil, BLACK)
	null.Left = null
	null.Right = null
	return null
}

func rotL(root *RBNode) *RBNode {
	x := root.Right
	root.Right = x.Left
	x.Left = root
	return root
}

func rotR(root *RBNode) *RBNode {
	x := root.Left
	root.Left = x.Right
	x.Right = root
	return root
}

func RBInsert(root *RBNode, val int) *RBNode {
	root = insertNode(root, val, 0)
	root.RB = BLACK
	return root
}
func insertNode(t *RBNode, val int, sw int) *RBNode {
	if t == nil {
		//每次新插入节点都是红色节点
		return NewRBNode(val, nil, nil, RED)
	}
	if t.Left.RB == RED && t.Right.RB == RED {
		t.RB = RED
		t.Left.RB = BLACK
		t.Right.RB = BLACK
	}
	if val < t.Val {
		t.Left = insertNode(t.Left, val, 0)
		if t.RB == RED && t.Left.RB == RED && sw == 1 {
			t = rotR(t)
		}
		if t.Left.RB == RED && t.Left.Left.RB == RED {
			t = rotR(t)
			t.RB = BLACK
			t.Right.RB = RED
		}
	} else {
		t.Right = insertNode(t.Right, val, 1)
		if t.RB == RED && t.Right.RB == RED && sw == 0 {
			t = rotL(t)
		}
		if t.Right.RB == RED && t.Right.Right.RB == RED {
			t = rotL(t)
			t.RB = BLACK
			t.Left.RB = RED
		}
	}
	return t
}
