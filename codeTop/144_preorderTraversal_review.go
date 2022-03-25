package codeTop

func preorderTraversalReview(root *TreeNode) []int {
	var ans []int
	return traversal144Review(root, ans)
}

func traversal144Review(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	if root.Left != nil {
		ans = traversal144Review(root.Left, ans)
	}
	if root.Right != nil {
		ans = traversal144Review(root.Right, ans)
	}
	return ans
}
