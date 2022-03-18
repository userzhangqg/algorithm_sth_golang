package codeTop

/*
二叉树遍历知识：https://www.jianshu.com/p/456af5480cee
？序遍历 指根节点的顺序

给定一个二叉树的根节点 root ，返回它的 中序 遍历。

完成￥￥￥

思路：
二叉树递归，中序：先递归左节点，根节点val append，再递归右节点
*/

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	ans = myInorderTraversal94(root, ans)
	return ans
}

func myInorderTraversal94(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	if root.Left != nil {
		ans = myInorderTraversal94(root.Left, ans)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = myInorderTraversal94(root.Right, ans)
	}
	return ans
}
