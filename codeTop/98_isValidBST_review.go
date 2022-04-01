package codeTop

import "math"

func isValidBSTReview(root *TreeNode) bool {
	return dfs98Review(root, math.MinInt64, math.MaxInt64)
}

func dfs98Review(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return dfs98Review(root.Left, lower, root.Val) && dfs98Review(root.Right, root.Val, upper)
}
