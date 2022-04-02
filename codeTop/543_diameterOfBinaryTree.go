package codeTop

/*
543. 二叉树的直径

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
这条路径可能穿过也可能不穿过根结点。

!!! 未做出
1. 深度优先维护最大值
*/

func diameterOfBinaryTree(root *TreeNode) int {
	var max int

	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		// 左儿子为根的子树的深度
		L := height(root.Left)
		// 右儿子为根的子树的深度
		R := height(root.Right)
		// 计算最长路径数
		max = max543(max, L+R)
		// 返回子树深度
		return max543(L, R) + 1
	}

	height(root)
	return max
}

func max543(x, y int) int {
	if x > y {
		return x
	}
	return y
}
