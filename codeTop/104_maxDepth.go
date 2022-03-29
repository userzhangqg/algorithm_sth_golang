package codeTop

/*
104. 二叉树的最大深度
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

给定一个二叉树，找出其最大深度
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明：叶子节点是指没有子节点的节点。

思路：
1. 层序遍历(广度优先）
2. 深度优先（递归）
*/

func maxDepth(root *TreeNode) int {
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
		//q = q[size:]
	}
	return ans
}
