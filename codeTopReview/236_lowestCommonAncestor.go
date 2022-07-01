package codeTopReview

/*
236. 二叉树的最近公共祖先
https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	rigth := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return rigth
	} else if rigth == nil {
		return left
	} else {
		return root
	}
}
