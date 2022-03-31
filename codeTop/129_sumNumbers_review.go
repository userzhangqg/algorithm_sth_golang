package codeTop

func sumNumbersReview(root *TreeNode) int {
	var ans int
	if root == nil {
		return ans
	}

	q := []*TreeNode{root}
	var tmp []int

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				tmp = append(tmp, node.Val)
			}

			if node.Left != nil {
				node.Left.Val = node.Val*10 + node.Left.Val
				q = append(q, node.Left)
			}

			if node.Right != nil {
				node.Right.Val = node.Val*10 + node.Right.Val
				q = append(q, node.Right)
			}
		}
	}

	for _, v := range tmp {
		ans += v
	}
	return ans
}
