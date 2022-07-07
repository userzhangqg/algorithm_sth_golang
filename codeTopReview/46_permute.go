package codeTopReview

/*
46. 全排列
https://leetcode.cn/problems/permutations/

因此时间复杂度为 O(n \times n!)O(n×n!)。

空间复杂度：O(n)O(n)，其中 nn 为序列的长度。除答案数组以外，递归函数在递归过程中需要为每一层递归函数分配栈空间，所以这里需要额外的空间且该空间取决于递归的深度，这里可知递归调用深度为 O(n)O(n)。
*/

func permute(nums []int) [][]int {
	var ans [][]int

	var backtrack func(nums []int, tmp []int)
	backtrack = func(nums []int, tmp []int) {
		if len(nums) == 0 {
			res := make([]int, len(tmp))
			copy(res, tmp)
			ans = append(ans, res)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 错误
			//backtrack(append(nums[:i], nums[i+1:]...), append(tmp, nums[i]))

			cur := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			tmp = append(tmp, cur)
			backtrack(nums, tmp)
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(nums, []int{})
	return ans
}
