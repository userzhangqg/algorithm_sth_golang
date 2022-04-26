package codeTop

func robReview(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	} else {
		if nums[0] > nums[1] {
			nums[1] = nums[0]
		}
	}

	for i := 2; i < size; i++ {
		if nums[i-1] > nums[i]+nums[i-2] {
			nums[i] = nums[i-1]
		} else {
			nums[i] = nums[i] + nums[i-2]
		}
	}
	return nums[size-1]
}
