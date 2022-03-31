package codeTop

/*
110. 平衡二叉树

给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一颗高度平衡二叉树定义为：
	一个二叉树每个节点的左右两个子树的高度差绝对值不超过1

!!! 未做出
思路：
从顶到底递归
*/

func isBalanced(root *TreeNode) bool {
	// 递归先设置返回条件
	if root == nil {
		return true
	}
	return abs110(height110(root.Left)-height110(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height110(root *TreeNode) int {
	// 计算最大高度
	if root == nil {
		return 0
	}
	return max110(height110(root.Left), height110(root.Right)) + 1
}

func max110(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs110(x int) int {
	// 计算绝对值
	if x < 0 {
		return -x
	}
	return x
}
