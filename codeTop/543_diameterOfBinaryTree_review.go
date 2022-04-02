package codeTop

func diameterOfBinaryTreeReview(root *TreeNode) int {
	var ans int

	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		L := height(root.Left)
		R := height(root.Right)

		ans = max543Review(ans, L+R)

		return max543Review(L, R) + 1
	}
	height(root)
	return ans
}

func max543Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}
