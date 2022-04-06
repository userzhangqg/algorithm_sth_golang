package codeTop

import "math"

/*
322. 零钱兑换
https://leetcode-cn.com/problems/coin-change/

给你一个整数数组coins，表示不同面额的硬币；以及一个整数amount，表示总金额。
计算并返回可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回-1

你可以认为每种硬币的数量是无限的。

!!! 做过忘记
1. 递归（常规递归超时）
2. 动态规划, 背包问题
https://leetcode-cn.com/problems/coin-change/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-wan-q-80r7/

*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min322(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min322(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
// 递归超出时间
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Slice(coins, func(i, j int) bool {
		if coins[i] > coins[j] {
			return true
		}
		return false
	})
	res := math.MaxInt64
	backtrack322(coins, 0, amount, 0, &res)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func backtrack322(coins []int, i int, amount int, total int, res *int) {
	if amount < 0 {
		return
	}
	if amount == 0 {
		*res = min322(*res, total)
	}

	for i < len(coins) {
		backtrack322(coins, i, amount-coins[i], total+1, res)
		i++
	}
}

func min322(x, y int) int {
	if x < y {
		return x
	}
	return y
}
*/
