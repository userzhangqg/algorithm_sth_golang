package codeTop

import "math"

func maxPathSum2(root *TreeNode) int {
	ans := math.MinInt32
	var maxSum func(*TreeNode) int
	maxSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max124(maxSum(node.Left), 0)
		rightMax := max124(maxSum(node.Right), 0)

		sum := node.Val + leftMax + rightMax
		ans = max124(ans, sum)

		return node.Val + max124(leftMax, rightMax)
	}
	maxSum(root)
	return ans
}

func max124(x, y int) int {
	if x > y {
		return x
	}
	return y
}
