package codeTop

/*
39. 组合总和
https://leetcode-cn.com/problems/combination-sum/
给你一个无重复元素的整数数组candidates和一个目标整数target，找出candidates中可以使数字和为目标数target的所有不同组合，并以列表形式返回。你可以按任意顺序返回这些组合。

candidates中的同一个数字可以无限制重复被选取。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为target的不同组合数少于150个。

$ Done
思路：
1. 回溯
*/

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	backtrack39(candidates, 0, target, []int{}, &ans)
	return ans
}

func backtrack39(nums []int, i int, target int, tmp []int, ans *[][]int) {
	if target <= 0 {
		if target == 0 {
			res := make([]int, len(tmp))
			copy(res, tmp)
			*ans = append(*ans, res)
		}
		return
	}

	for i < len(nums) {
		backtrack39(nums, i, target-nums[i], append(tmp, nums[i]), ans)
		i++
	}
}
