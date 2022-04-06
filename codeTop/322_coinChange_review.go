package codeTop

import "math"

func coinChangeReview(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < len(dp); j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min322Review(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min322Review(x, y int) int {
	if x < y {
		return x
	}
	return y
}
