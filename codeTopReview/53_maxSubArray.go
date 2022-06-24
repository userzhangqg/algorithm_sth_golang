package codeTopReview

/*
53. 最大连续子数组和
https://leetcode.cn/problems/maximum-subarray/

思路：
1. 动态规划
*/

func maxSubArray(nums []int) int {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

//func maxSubArray(nums []int) int {
//	ans := nums[0]
//	cur := ans
//
//	for i := 1; i < len(nums); i++ {
//		// "如果我们在一起会变得更好，那我们就在一起。否则，我就丢下你" --- 一别两宽，各生欢喜
//		if cur+nums[i] > nums[i] {
//			cur += nums[i]
//		} else {
//			cur = nums[i]
//		}
//
//		if cur > ans {
//			ans = cur
//		}
//	}
//	return ans
//}
