package codeTop

/*
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

「代码随想录」带你学透DP子序列问题！718. 最长重复子数组:【动态规划】详解
https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/solution/dai-ma-sui-xiang-lu-718-zui-chang-zhong-rowbh/
*/

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var ans int
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
