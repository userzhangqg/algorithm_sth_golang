package codeTop

import "math"

/*
98. 验证二叉搜索树

给你一个二叉树的根节点root，判断其是否是一个有效的二叉搜索树。
有效的二叉搜索树定义如下：
- 节点的左子树只包含小于当前节点的数
- 节点的右子树只包含大于当前节点的数
- 所有左子树和右子树自身必须也是二叉搜索树
*/

func isValidBST(root *TreeNode) bool {
	return dfs98(root, math.MinInt64, math.MaxInt64)
}

func dfs98(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return dfs98(root.Left, lower, root.Val) && dfs98(root.Right, root.Val, upper)
}
