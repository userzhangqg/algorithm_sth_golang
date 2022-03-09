package codeTop

/*
102. 二叉树的层序遍历
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
给你二叉树的根节点root，返回其节点值的层序遍历。（即逐层地，从左到右访问所有节点）。
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/die-dai-di-gui-duo-tu-yan-shi-102er-cha-shu-de-cen/

!!! 未做出
正确思路：
1. 广度优先
2. 深度优先递归
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		var tmp []int
		size := len(q)

		for i := 0; i < size; i++ {
			// 注释的方式空间复杂度低??? 待学习
			//node := q[0]
			//q = append(q[:0], q[1:]...)
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
		res = append(res, tmp)
	}
	return res
}
