package codeTop

func sortArray912MergeReview(nums []int) []int {
	var merge func(left, right []int) []int
	merge = func(left, right []int) []int {
		l, r, i := 0, 0, 0
		newNums := make([]int, len(left)+len(right))

		for l < len(left) && r < len(right) {
			if left[l] < right[r] {
				newNums[i] = left[l]
				l++
			} else {
				newNums[i] = right[r]
				r++
			}
			i++
		}
		copy(newNums[i:], left[l:])
		copy(newNums[i+len(left)-l:], right[r:])
		return newNums
	}

	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		// 递归跳出
		if len(nums) <= 1 {
			return nums
		}
		mid := len(nums) / 2

		left := sort(nums[:mid])
		right := sort(nums[mid:])

		return merge(left, right)
	}
	return sort(nums)
}
