package basic_algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
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
func inorderTraversal(root *TreeNode) []int {
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
