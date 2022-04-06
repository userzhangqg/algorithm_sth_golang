package codeTop

func minPathSumReview(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for m := 1; m < len(grid); m++ {
		dp[m][0] = dp[m-1][0] + grid[m][0]
	}
	for n := 1; n < len(grid[0]); n++ {
		dp[0][n] = dp[0][n-1] + grid[0][n]
	}

	for m := 1; m < len(grid); m++ {
		for n := 1; n < len(grid[0]); n++ {
			dp[m][n] = min64Review(dp[m-1][n], dp[m][n-1]) + grid[m][n]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min64Review(x, y int) int {
	if x < y {
		return x
	}
	return y
}
