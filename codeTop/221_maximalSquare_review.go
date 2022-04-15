package codeTop

func maximalSquareReview(matrix [][]byte) int {
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
					dp[i][j] = min221Review(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				}

				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func min221Review(args ...int) int {
	min := args[0]

	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}
