package codeTop

func buildTree105(preorder []int, inorder []int) *TreeNode {
	return buildTreeReview(preorder, inorder)
}

func buildTreeReview(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	i := 0

	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = buildTreeReview(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTreeReview(preorder[i+1:], inorder[i+1:])

	return root
}
