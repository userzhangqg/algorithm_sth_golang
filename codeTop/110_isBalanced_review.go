package codeTop

func isBalancedReview(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs110Review(height110Review(root.Left)-height110Review(root.Right)) <= 1 &&
		isBalancedReview(root.Left) && isBalancedReview(root.Right)
}

func height110Review(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max110Review(height110Review(root.Left), height110Review(root.Right)) + 1
}

func abs110Review(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max110Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}
