package codeTop

func inorderTraversalReview(root *TreeNode) []int {
	return inorderReview(root, make([]int, 0))
}

func inorderReview(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}

	if root.Left != nil {
		ans = inorderReview(root.Left, ans)
	}

	ans = append(ans, root.Val)

	if root.Right != nil {
		ans = inorderReview(root.Right, ans)
	}
	return ans
}
