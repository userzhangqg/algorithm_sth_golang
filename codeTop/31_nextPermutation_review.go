package codeTop

func nextPermutationReview(nums []int) {
	size := len(nums)

	i := size - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 有可能全部降序，i为-1
	if i >= 0 {
		j := size - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse31Review(nums[i+1:])
}

func reverse31Review(nums []int) {
	size := len(nums)

	for i := 0; i < size/2; i++ {
		nums[i], nums[size-1-i] = nums[size-1-i], nums[i]
	}
}
