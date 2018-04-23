package basic_algorithm

import (
	"io"
	"fmt"
)

const (
	BLACK = iota
	RED
)

var null *RBNode

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

func rotL(root *RBNode) *RBNode {
	x := root.Right
	root.Right = x.Left
	x.Left = root
	return x
}

func rotR(root *RBNode) *RBNode {
	x := root.Left
	root.Left = x.Right
	x.Right = root
	return x
}

func RBInsert(root *RBNode, val int) *RBNode {
	root = insertNode(root, val, 0)
	root.RB = BLACK
	return root
}
func insertNode(t *RBNode, val int, sw int) *RBNode {
	if t == null {
		//每次新插入节点都是红色节点
		return NewRBNode(val, null, null, RED)
	}
	//4-节点 分裂
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
