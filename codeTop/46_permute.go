package codeTop

/*
46. 全排列
https://leetcode-cn.com/problems/permutations/

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

!!! 忘记
思路：回溯，未遍历+已遍历
*/

var ans [][]int

func permute(nums []int) [][]int {
	findAll46(nums, make([]int, 0))
	return ans
}

func findAll46(nums []int, tmp []int) {
	if len(nums) == 0 {
		slice := make([]int, len(tmp))
		copy(slice, tmp)
		ans = append(ans, slice)
		return
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		tmp = append(tmp, cur)
		nums = append(nums[:i], nums[i+1:]...)
		findAll46(nums, tmp)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		tmp = tmp[:len(tmp)-1]
	}
}
