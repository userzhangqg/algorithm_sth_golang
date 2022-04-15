package codeTop

/*
221. 最大正方形
https://leetcode-cn.com/problems/maximal-square/
在一个由'0' 和 '1' 组成的二维矩阵内，找到只包含'1'的最大正方形，并返回其面积

!!! TODO
思路：
动态规划：
dp[i][j] 代表以i，j为右下角的正方形最大边长
*/

func maximalSquare(matrix [][]byte) int {
	rows, columns := len(matrix), len(matrix[0])

	dp := make([][]int, rows)

	var maxSide int

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min221(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}

				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func min221(args ...int) int {
	min := args[0]

	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}
