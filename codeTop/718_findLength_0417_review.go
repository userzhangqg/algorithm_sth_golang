package codeTop

func findLength0417Review(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	var ans int
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

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
