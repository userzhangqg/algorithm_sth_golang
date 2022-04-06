package codeTop

/*
78. 子集
https://leetcode-cn.com/problems/subsets/
给你一个整数数组nums，数组中的元素互不相同。返回该数组所有可能的子集（幂集）
解集 不能 包含重复的子集。你可以按任意顺序返回解集。

!!! TODO 做出待复习
思路：
1. 回溯（与官方答案顺序不一致）
*/

func subsets(nums []int) [][]int {
	var ans [][]int
	backtrack78(nums, 0, []int{}, &ans)
	return ans
}

func backtrack78(nums []int, i int, tmp []int, ans *[][]int) {
	res := make([]int, len(tmp))
	copy(res, tmp)
	*ans = append(*ans, res)

	for i < len(nums) {
		backtrack78(nums, i+1, append(tmp, nums[i]), ans)
		i++
	}
}
