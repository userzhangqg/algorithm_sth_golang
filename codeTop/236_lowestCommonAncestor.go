package codeTop

/*
236. 二叉树的最近公共祖先

给定一个二叉树，找到该树中两个指定节点的最近公共祖先

百度百科：最近公共祖先的定义为：对于有根树T的两个节点p、q，最近公共祖先表示为一个节点x，满足x是p、q的祖先且x的深度尽可能大（一个节点也可以是它自己的祖先）

!!! 未做出

思路：二叉树解法优先想到递归，设置跳出条件
https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//if left != nil && right != nil {
	//	return root
	//}
	//
	//if left == nil {
	//	return right
	//}else {
	//	return left
	//}
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}
