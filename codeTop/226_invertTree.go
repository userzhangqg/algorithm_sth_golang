package codeTop

/*
226. 翻转二叉树

给你一棵二叉树的根节点root，翻转这棵二叉树，并返回其根节点。

*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
