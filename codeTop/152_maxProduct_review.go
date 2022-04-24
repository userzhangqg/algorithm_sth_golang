package codeTop

import "math"

func maxProductReview(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max152Review(imax*nums[i], nums[i])
		imin = min152Review(imin*nums[i], nums[i])
		max = max152Review(max, imax)
	}
	return max
}

func max152Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min152Review(x, y int) int {
	if x < y {
		return x
	}
	return y
}
