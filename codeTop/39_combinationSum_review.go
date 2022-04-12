package codeTop

func combinationSumReview(candidates []int, target int) [][]int {
	var ans [][]int
	backtrack39Review(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack39Review(nums []int, target int, i int, tmp []int, ans *[][]int) {
	if target <= 0 {
		if target == 0 {
			res := make([]int, len(tmp))
			copy(res, tmp)
			*ans = append(*ans, res)
		}
		return
	}

	for i < len(nums) {
		backtrack39Review(nums, target-nums[i], i, append(tmp, nums[i]), ans)
		i++
	}
}
