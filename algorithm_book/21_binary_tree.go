package algorithm_book

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(r, root.Val)
	r = append(r, PreOrderTree(root.Left)...)
	r = append(r, PreOrderTree(root.Right)...)
	return r
}
func InOrderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	r = append(r, InOrderTree(root.Left)...)
	r = append(r, root.Val)
	r = append(r, InOrderTree(root.Right)...)
	return r
}

func PreOrderTreeStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) != 0 {
		//注意 := 的深坑！！！
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

//TODO
func InOrderTreeStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	result := []int{}

	//左 中 右
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, root.Val)
			root = root.Right
		}

	}
	return result
}

//先序 根->左->右
//后序 左->右->根
//根->右->左  反转得到 左->右->根
func PostOrderTreeStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	result := []int{}

	stack = append(stack, root)

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		result = append(result, cur.Val)
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	mid := len(result) / 2
	for i := 0; i < mid; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
//层次遍历二叉树
func BinaryTreeLevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	result := [][]int{}
	for len(queue) != 0 {
		tmp := []*TreeNode{}
		level := []int{}
		for i := 0; i < len(queue); i++ {
			cur := queue[i]
			level = append(level, cur.Val)
			if cur.Left != nil {
				tmp = append(tmp, cur.Left)
			}
			if cur.Right != nil {
				tmp = append(tmp, cur.Right)
			}
		}
		result = append(result, level)
		queue = tmp
	}
	return result
}
