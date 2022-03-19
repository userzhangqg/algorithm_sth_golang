package codeTop

func climbStairsReview(n int) int {
	x, y := 1, 1

	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}
