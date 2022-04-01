package codeTop

/*
113. 路径总和II
https://leetcode-cn.com/problems/path-sum-ii/
给你一个二叉树的根节点root和一个整数目标和targetSum，找出所有从根节点到叶子节点 路径总和等于给定目标和的路径

叶子节点 是指没有子节点的节点

$$$ Done
思路：
回溯
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	backtrack(root, targetSum, 0, []int{}, &ans)
	return ans
}

func backtrack(root *TreeNode, target int, sum int, tmp []int, ans *[][]int) {
	if root == nil {
		return
	}
	sum += root.Val
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && target == sum {
		// 需使用make定义长度
		res := make([]int, len(tmp))
		copy(res, tmp)
		*ans = append(*ans, res)
		return
	}

	backtrack(root.Left, target, sum, tmp, ans)
	backtrack(root.Right, target, sum, tmp, ans)
}
