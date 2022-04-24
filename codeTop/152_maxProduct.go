package codeTop

import "math"

/*
152. 乘积最大子数组

给你一个整数数组nums，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积

子数组 是数组的连续子序列

!!! TODO
思路：
1. 动态规划
https://leetcode-cn.com/problems/maximum-product-subarray/solution/hua-jie-suan-fa-152-cheng-ji-zui-da-zi-xu-lie-by-g/
*/

func maxProduct(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max152(imax*nums[i], nums[i])
		imin = min152(imin*nums[i], nums[i])
		max = max152(max, imax)
	}
	return max
}

func min152(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max152(x, y int) int {
	if x > y {
		return x
	}
	return y
}
