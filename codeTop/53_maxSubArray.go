package codeTop

/*
https://leetcode-cn.com/problems/maximum-subarray/
53. 最大子数组和

给你一个整数数组nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素）返回其最大和

子数组是数组中的一个连续部分。

无思路=:=

正确思路：
1. 动态规划
2. 分治算法（待学习）
我觉得这道题目的思想是： 走完这一生 如果我和你在一起会变得更好，那我们就在一起，否则我就丢下你。 我回顾我最光辉的时刻就是和不同人在一起，变得更好的最长连续时刻
*/

func maxSubArray(nums []int) int {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > ans {
			ans = nums[i]
		}
	}

	return ans
}
