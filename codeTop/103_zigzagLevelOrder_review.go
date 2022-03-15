package codeTop

func zigzagLevelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	level := 1

	for len(q) > 0 {
		tmp := []int{}

		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = append(q[:0], q[1:]...)
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if level%2 == 0 {
			for m, n := 0, len(tmp)-1; m < n; m, n = m+1, n-1 {
				tmp[m], tmp[n] = tmp[n], tmp[m]
			}
		}
		level++
		res = append(res, tmp)
	}
	return res
}
