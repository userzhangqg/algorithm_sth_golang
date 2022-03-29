package codeTop

func maxDepthReview(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans++
	}
	return ans
}
