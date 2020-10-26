package binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var stack []*TreeNode
	var ret []int
	stack = append(stack, root)

	for len(stack) != 0 {
		e := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ret = append(ret, e.Val)

		if e.Right != nil {
			stack = append(stack, e.Right)
		}

		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}

	return ret
}