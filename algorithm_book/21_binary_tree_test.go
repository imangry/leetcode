package algorithm_book

import (
	"testing"
	"fmt"
)

func TestPreOrderTreeStack(t *testing.T) {
	left := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	right := &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{Left: left, Right: right, Val: 0}
	fmt.Println(PreOrderTreeStack(root))
}
