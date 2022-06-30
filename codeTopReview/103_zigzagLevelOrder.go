package codeTopReview

/*
103. 二叉树的锯齿形层序遍历
https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	q := []*TreeNode{root}

	for level := 1; len(q) > 0; level++ {
		size := len(q)
		var tmp []int

		for i := 0; i < size; i++ {
			node := q[i]

			tmp = append(tmp, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		q = q[size:]

		if level%2 == 0 {
			// 前面和层序遍历一样，需翻转偶数层
			for m, n := 0, len(tmp)-1; m < n; m, n = m+1, n-1 {
				tmp[m], tmp[n] = tmp[n], tmp[m]
			}
		}

		ans = append(ans, tmp)
	}
	return ans
}
