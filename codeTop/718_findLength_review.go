package codeTop

/*
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

*/

func findLengthReview(nums1 []int, nums2 []int) int {
	var ans int
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ans = max718Review(ans, dp[i][j])
		}
	}
	return ans
}

func max718Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}
