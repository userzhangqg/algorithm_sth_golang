package codeTop

func hasPathSumReview(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSumReview(root.Left, targetSum) || hasPathSumReview(root.Right, targetSum)
}
