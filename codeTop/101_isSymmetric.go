package codeTop

/*
101. 对称二叉树
给你一个二叉树的根节点root，检查它是否轴对称

$$$ Done
思路：
递归
*/

func isSymmetric(root *TreeNode) bool {
	return isTrue101(root.Left, root.Right)
}

func isTrue101(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}

	return isTrue101(node1.Left, node2.Right) && isTrue101(node1.Right, node2.Left)
}
