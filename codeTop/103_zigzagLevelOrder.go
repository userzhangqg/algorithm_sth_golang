package codeTop

/*
103. 二叉树的锯齿形层序遍历
https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

给你二叉树的根节点root，返回其节点值的锯齿形层序遍历。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）

！！！ 思路错误
正确思路：
1. 翻转数组
2. 使用标记 https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/solution/go-shuang-bai-by-ba-xiang-3/
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	q := []*TreeNode{root}

	for level := 0; len(q) > 0; level++ {
		var tmp []int

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

		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for m, n := 0, len(tmp)-1; m < n; m, n = m+1, n-1 {
				tmp[m], tmp[n] = tmp[n], tmp[m]
			}
		}

		ans = append(ans, tmp)
	}
	return ans
}
