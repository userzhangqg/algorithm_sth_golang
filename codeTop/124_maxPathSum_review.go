package codeTop

func maxPathSumReview(root *TreeNode) int {
	var ans = -1001

	var leftOrRightMax func(*TreeNode) int
	leftOrRightMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := max124Review(leftOrRightMax(node.Left), 0)
		rightMax := max124Review(leftOrRightMax(node.Right), 0)

		sum := node.Val + leftMax + rightMax

		ans = max124Review(sum, ans)

		return node.Val + max124Review(leftMax, rightMax)
	}
	leftOrRightMax(root)
	return ans
}

func max124Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}
