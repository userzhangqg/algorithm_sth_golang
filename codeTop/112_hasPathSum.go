package codeTop

/*
112. 路径总和

给你二叉树的根节点root和一个目标和的整数targetSum。判断该树中是否存在根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum。
如果存在，返回true；否则，返回false。
叶子节点 是指没有子节点的节点。
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	t := targetSum - root.Val
	if root.Left == nil && root.Right == nil {
		return t == 0
	}

	return hasPathSum(root.Left, t) || hasPathSum(root.Right, t)
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	// 繁琐做法
	if root == nil {
		return false
	}
	return isTrue112(root, targetSum, 0)
}

func isTrue112(root *TreeNode, targetSum int, curSum int) bool {
	if root == nil && curSum == targetSum {
		return true
	}
	if root == nil {
		return false
	}

	curSum += root.Val

	if root.Left == nil {
		return isTrue112(root.Right, targetSum, curSum)
	}
	if root.Right == nil {
		return isTrue112(root.Left, targetSum, curSum)
	}
	return isTrue112(root.Right, targetSum, curSum) || isTrue112(root.Left, targetSum, curSum)
}
