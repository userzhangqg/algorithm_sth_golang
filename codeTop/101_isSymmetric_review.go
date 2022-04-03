package codeTop

func isSymmetricReview(root *TreeNode) bool {
	return check101(root.Left, root.Right)
}

func check101(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && check101(node1.Left, node2.Right) && check101(node1.Right, node2.Left)
}
