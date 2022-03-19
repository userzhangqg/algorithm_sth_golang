package codeTop

func rightSideViewReview(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	var cent [][]int

	q := []*TreeNode{root}

	for len(q) > 0 {
		var tmp []int
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[i]

			tmp = append(tmp, node.Val)
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
		}
		q = q[size:]
		cent = append(cent, tmp)
	}

	for _, v := range cent {
		ans = append(ans, v[0])
	}
	return ans
}
