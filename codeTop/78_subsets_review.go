package codeTop

func subsetsReview(nums []int) [][]int {
	var ans [][]int

	backtrack78Review(nums, 0, []int{}, &ans)
	return ans
}

func backtrack78Review(nums []int, i int, tmp []int, ans *[][]int) {
	res := make([]int, len(tmp))
	copy(res, tmp)
	*ans = append(*ans, res)

	for i < len(nums) {
		backtrack78Review(nums, i+1, append(tmp, nums[i]), ans)
		i++
	}
}
