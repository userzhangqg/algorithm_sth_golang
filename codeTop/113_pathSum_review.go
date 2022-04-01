package codeTop

func pathSumReview(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	backtrack113Review(root, &ans, []int{}, targetSum, 0)
	return ans
}

func backtrack113Review(root *TreeNode, ans *[][]int, tmp []int, targetSum int, curSum int) {
	if root == nil {
		return
	}
	tmp = append(tmp, root.Val)
	curSum += root.Val
	if root.Left == nil && root.Right == nil && curSum == targetSum {
		res := make([]int, len(tmp))
		copy(res, tmp)
		*ans = append(*ans, res)
	}

	backtrack113Review(root.Left, ans, tmp, targetSum, curSum)
	backtrack113Review(root.Right, ans, tmp, targetSum, curSum)
}
