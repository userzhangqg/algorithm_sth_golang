package codeTop

/*
199. 二叉树的右视图
https://leetcode-cn.com/problems/binary-tree-right-side-view/

给定一个二叉树的根节点root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

!!! 细节错误，二叉树层序遍历还需学习

思路：
https://leetcode-cn.com/problems/binary-tree-right-side-view/solution/dai-ma-sui-xiang-lu-wo-yao-da-shi-ge-er-mdkms/
层序遍历，取每层最后一个元素
*/

func rightSideView(root *TreeNode) []int {
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
