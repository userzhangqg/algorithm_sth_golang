package codeTopReview

/*
199. 二叉树的右视图
https://leetcode.cn/problems/binary-tree-right-side-view/

复杂度分析
时间复杂度 :O(n)。深度优先搜索最多访问每个结点一次，因此是线性复杂度。
空间复杂度 : O(n)。最坏情况下，栈内会包含接近树高度的结点数量，占用O(n) 的空间。
*/

func rightSideView(root *TreeNode) []int {
	var ans []int

	if root == nil {
		return ans
	}

	var res [][]int
	q := []*TreeNode{root}

	for len(q) > 0 {
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

		res = append(res, tmp)
		q = q[size:]
	}

	for _, v := range res {
		ans = append(ans, v[len(v)-1])
	}
	return ans
}
