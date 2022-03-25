package codeTop

/*
144. 二叉树的前序遍历

给你一个二叉树的根节点root，返回它节点值的前序遍历

进阶：你可以使用迭代方式吗

*/

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	ans = traversal144(root, ans)
	return ans
}

func traversal144(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)

	ans = traversal144(root.Left, ans)
	ans = traversal144(root.Right, ans)
	return ans
}
