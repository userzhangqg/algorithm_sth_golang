package codeTop

/*
64. 最小路径和
https://leetcode-cn.com/problems/minimum-path-sum/
给定一个包含非负整数的m x n网格grid，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步

$$$ Done
思路：
1. 动态规划dp
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for m := 1; m < len(grid); m++ {
		dp[m][0] = grid[m][0] + dp[m-1][0]
	}
	for n := 1; n < len(grid[0]); n++ {
		dp[0][n] = grid[0][n] + dp[0][n-1]
	}

	for m := 1; m < len(grid); m++ {
		for n := 1; n < len(grid[0]); n++ {
			dp[m][n] = min64(dp[m-1][n], dp[m][n-1]) + grid[m][n]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min64(x, y int) int {
	if x < y {
		return x
	}
	return y
}
