package codeTopReview

/*
72. 编辑距离
https://leetcode.cn/problems/edit-distance/

复杂度分析

时间复杂度 ：
O(mn)，其中
m 为 word1 的长度，
n 为 word2 的长度。
空间复杂度 ：
O(mn)，我们需要大小为
O(mn) 的
D 数组来记录状态值。

*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)

	for d := range dp {
		dp[d] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min72(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min72(args ...int) int {
	ans := args[0]

	for _, v := range args {
		if v < ans {
			ans = v
		}
	}
	return ans
}
