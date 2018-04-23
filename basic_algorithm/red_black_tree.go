package basic_algorithm

import (
	"io"
	"fmt"
)

const (
	BLACK = iota
	RED
)
const (
	LEFT  = iota
	RIGHT
)

var null *RBNode

//空节点 都是黑色节点
func RBInit() *RBNode {
	null = NewRBNode(0, nil, nil, BLACK)
	null.Left = null
	null.Right = null
	return null
}

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

func PrintRBTree(root *RBNode, writer io.Writer) {
	if root == null {
		return
	}
	stack := []*RBNode{}
	stack = append(stack, root)
	var first *RBNode
	for len(stack) != 0 {
		buf := []*RBNode{}
		for len(stack) != 0 {
			first, stack = stack[0], stack[1:]
			var red string
			if first.RB == RED {
				red = "+"
			}
			fmt.Fprint(writer, first.Val, red, "\t")
			if first.Left != null {
				buf = append(buf, first.Left)
			}
			if first.Right != null {
				buf = append(buf, first.Right)
			}
		}
		fmt.Fprintln(writer)
		stack = append(stack, buf...)
	}
}

//当前节点左旋
func rotL(root *RBNode) *RBNode {
	x := root.Right
	root.Right = x.Left
	x.Left = root
	return x
}

//当前节点右旋
func rotR(root *RBNode) *RBNode {
	x := root.Left
	root.Left = x.Right
	x.Right = root
	return x
}

func RBInsert(root *RBNode, val int) *RBNode {
	root = insertNode(root, val, LEFT)
	root.RB = BLACK
	return root
}
func insertNode(t *RBNode, val int, sw int) *RBNode {
	if t == null {
		//每次新插入节点都是红色节点
		return NewRBNode(val, null, null, RED)
	}
	//在插入节点前，将所有的4-节点分裂
	//4-节点 分裂
	if t.Left.RB == RED && t.Right.RB == RED {
		t.RB = RED
		t.Left.RB = BLACK
		t.Right.RB = BLACK
	}
	if val < t.Val {
		t.Left = insertNode(t.Left, val, LEFT)
		//如果在搜索路径上有两个在不同方向上的红链接，则从下面的节点进行一次旋转；简化成下一步处理的另一种情况
		if t.RB == RED && t.Left.RB == RED && sw == RIGHT {
			t = rotR(t)
		}
		//如果在搜索路径上有两个在相同方向上的红链接，则从上面的节点进行一次旋转
		if t.Left.RB == RED && t.Left.Left.RB == RED {
			t = rotR(t)
			t.RB = BLACK
			t.Right.RB = RED
		}
	} else {
		t.Right = insertNode(t.Right, val, RIGHT)
		if t.RB == RED && t.Right.RB == RED && sw == LEFT {
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
