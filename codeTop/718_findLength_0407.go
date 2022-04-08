package codeTop

/*
718. 最长重复子数组

给两个整数数组nums1 和 nums2，返回两个数组中公共的、长度最长的子数组的长度。

*/

func findLength0407(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	var ans int
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ans = max0407(ans, dp[i][j])
		}
	}
	return ans
}

func max0407(x, y int) int {
	if x > y {
		return x
	}
	return y
}
