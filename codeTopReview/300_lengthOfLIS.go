package codeTopReview

/*
300.最长递增子序列
https://leetcode.cn/problems/longest-increasing-subsequence/

思路：
1. 动态规划
复杂度分析

时间复杂度：O(n^2)O(n
2
 )，其中 nn 为数组 \textit{nums}nums 的长度。动态规划的状态数为 nn，计算状态 dp[i]dp[i] 时，需要 O(n)O(n) 的时间遍历 dp[0 \ldots i-1]dp[0…i−1] 的所有状态，所以总时间复杂度为 O(n^2)O(n
2
 )。

空间复杂度：O(n)O(n)，需要额外使用长度为 nn 的 dpdp 数组。

*/

func lengthOfLIS(nums []int) int {
	// 定义 dp[i] 为考虑前 ii 个元素，以第 ii 个数字结尾的最长上升子序列的长度，注意 nums[i] 必须被选取。

	var dp []int
	var ans int

	for i, v := range nums {
		dp = append(dp, 1)

		for j := 0; j < i; j++ {
			if v > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
