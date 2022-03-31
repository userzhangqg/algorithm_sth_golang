package codeTop

/*
129. 求根节点到叶节点数字之和
https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

给你一个二叉树的根节点root，树中每个节点都存放有一个0到9之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如：从根节点到叶子节点的路径 1 -> 2 -> 3 表示数字123.
计算从根节点到叶子节点生成的所有数字之和。

叶节点是指没有子节点的节点。

做出,>^<
思路:
层序遍历，判断是否无叶结点。
*/

func sumNumbers(root *TreeNode) int {
	var ans int
	var tmp []int

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				tmp = append(tmp, node.Val)
				continue
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
