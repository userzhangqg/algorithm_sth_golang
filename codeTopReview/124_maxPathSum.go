package codeTopReview

import "math"

/*
124. 二叉树中的最大路径和

复杂度分析

时间复杂度：O(N)O(N)，其中 NN 是二叉树中的节点个数。对每个节点访问不超过 22 次。

空间复杂度：O(N)O(N)，其中 NN 是二叉树中的节点个数。空间复杂度主要取决于递归调用层数，最大层数等于二叉树的高度，最坏情况下，二叉树的高度等于二叉树中的节点个数。

*/

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var maxSum func(root *TreeNode) int
	maxSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftMax := max124(maxSum(root.Left), 0)
		rightMax := max124(maxSum(root.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		sum := root.Val + leftMax + rightMax

		// 更新答案
		ans = max124(ans, sum)

		// 返回节点的最大贡献值
		return root.Val + max124(leftMax, rightMax)
	}
	maxSum(root)
	return ans
}

func max124(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
