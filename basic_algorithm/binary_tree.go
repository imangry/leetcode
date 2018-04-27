package basic_algorithm

import (
	"io"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func InsertTreeNode(root *TreeNode, val int) *TreeNode {
	//新建节点
	if root == nil {
		return NewTreeNode(val, nil, nil)
	}
	//将节点挂载到左右子树上
	if val < root.Val {
		root.Left = InsertTreeNode(root.Left, val)
	} else {
		root.Right = InsertTreeNode(root.Right, val)
	}
	return root
}

func RemoveTreeNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = RemoveTreeNode(root.Left, val)
	} else if val > root.Val {
		root.Right = RemoveTreeNode(root.Right, val)
	} else {
		var x *TreeNode
		//先查左子树，找到最大的节点，替换掉当前值；再将最大节点删掉
		if root.Left != nil {
			for x = root.Left; x.Right != nil; x = x.Right {
			}
			root.Val = x.Val
			root.Left = RemoveTreeNode(root.Left, x.Val)
			//再查右子树，找到最小的节点，替换当前值；再将最小节点删掉
		} else if root.Right != nil {
			for x = root.Right; x.Left != nil; x = x.Left {
			}
			root.Val = x.Val
			root.Right = RemoveTreeNode(root.Right, x.Val)
			//当前节点没有子树，直接将当前节点置为nil
		} else {
			root = nil
		}
	}
	return root
}

func PrintTree(root *TreeNode, w io.Writer) {
	w.Write([]byte("("))
	if root != nil {
		w.Write([]byte(strconv.Itoa(root.Val)))
		PrintTree(root.Left, w)
		PrintTree(root.Right, w)
	}
	w.Write([]byte(")"))
}

func SearchTree(root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val < val {
		return SearchTree(root.Left, val)
	} else if root.Val > val {
		return SearchTree(root.Right, val)
	} else {
		return true
	}
}

func PreOrderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, p.Val)

		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return result
}
func InOrderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}

	p := root
	for len(stack) != 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result = append(result, p.Val)
			p = p.Right
		}
	}

	return result
}

func BFS(head *TreeNode) []int {
	result := []int{}

	if head == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, head)
	var first *TreeNode
	for len(queue) > 0 {
		first, queue = queue[0], queue[1:]
		result = append(result, first.Val)

		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
	}
	return result
}
