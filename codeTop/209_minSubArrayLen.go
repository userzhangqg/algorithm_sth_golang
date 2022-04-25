package codeTop

import "math"

/*
209. 长度最小的子数组

给定一个含有n个正整数的数组和一个正整数target
找出该数组中满足其和>= target 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的子数组，返回0

!!! TODO
做过，思路不清
思路：
1. 滑动窗口
*/

func minSubArrayLen(target int, nums []int) int {
	l, r := -1, 0
	var sum int
	minSize := math.MaxInt64
	for r < len(nums) {
		if nums[r] == target {
			return 1
		}
		sum += nums[r]
		tmp := sum
		for l < r && tmp > target && tmp-nums[l+1] >= target {
			tmp -= nums[l+1]
			l++
		}
		sum = tmp
		if sum >= target && r-l < minSize {
			minSize = r - l
		}
		r++
	}
	if minSize == math.MaxInt64 {
		return 0
	}
	return minSize
}
