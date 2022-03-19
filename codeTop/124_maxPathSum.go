package codeTop

/*
124. 二叉树中的最大路径和

路径 被定义为一条从树中任意节点出发，沿着父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中至多出现一次。该路径至少包含一个节点，且不一定经过根节点。
路径和是路径中各节点值的总和。
给你一个二叉树的根节点root，返回其最大路径和
*/

var ans124 int

func maxPathSum(root *TreeNode) int {
	ans124 = -1001
	dfs124(root)
	return ans124
}

func dfs124(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归计算左右子节点的最大贡献值
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	leftMax := getMax(dfs124(root.Left), 0)
	rightMax := getMax(dfs124(root.Right), 0)

	// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
	sum := root.Val + leftMax + rightMax
	// 更新答案
	ans124 = getMax(ans124, sum)

	return root.Val + getMax(leftMax, rightMax)
}

func getMax(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
