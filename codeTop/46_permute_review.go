package codeTop

func permuteReview(nums []int) [][]int {
	var ans [][]int
	ans = findAll46Review(nums, make([]int, 0), ans)
	return ans
}

func findAll46Review(nums []int, tmp []int, ans [][]int) [][]int {
	if len(nums) == 0 {
		slice := make([]int, len(tmp))
		copy(slice, tmp)
		ans = append(ans, slice)
		return ans
	}

	for i := 0; i < len(nums); i++ {
		t := nums[i]
		nums = append(nums[:i], nums[i+1:]...)
		tmp = append(tmp, t)
		ans = findAll46Review(nums, tmp, ans)
		nums = append(nums[:i], append([]int{t}, nums[i:]...)...)
		tmp = tmp[:len(tmp)-1]
	}
	return ans
}
