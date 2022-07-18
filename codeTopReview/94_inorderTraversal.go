package codeTopReview

/*
94. 二叉树的中序遍历
https://leetcode.cn/problems/binary-tree-inorder-traversal/

复杂度分析

时间复杂度：O(n)，其中 n 为二叉树节点的个数。二叉树的遍历中每个节点会被访问一次且只会被访问一次。
空间复杂度：
O(n)。空间复杂度取决于递归的栈深度，而栈深度在二叉树为一条链的情况下会达到O(n) 的级别。

*/

func inorderTraversal(root *TreeNode) []int {
	var ans []int

	var traver func(root *TreeNode)
	traver = func(root *TreeNode) {
		if root == nil {
			return
		}

		traver(root.Left)
		ans = append(ans, root.Val)
		traver(root.Right)
	}
	traver(root)
	return ans
}
