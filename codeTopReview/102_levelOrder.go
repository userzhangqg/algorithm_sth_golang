package codeTopReview

/*
102. 二叉树层序遍历
https://leetcode.cn/problems/binary-tree-level-order-traversal/
*/

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	q := []*TreeNode{root}

	for len(q) != 0 {
		var tmp []int

		size := len(q)

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

		ans = append(ans, tmp)
		q = q[size:]
	}
	return ans
}
