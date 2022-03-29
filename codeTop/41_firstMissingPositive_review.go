package codeTop

func firstMissingPositiveReview(nums []int) int {
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			nums[i] = size + 1
		}
	}

	for i := 0; i < size; i++ {
		num := abs41Review(nums[i])
		if num < size+1 {
			nums[num-1] = -abs41Review(nums[num-1])
		}
	}

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return size + 1
}

func abs41Review(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
