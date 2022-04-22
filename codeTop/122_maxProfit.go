package codeTop

/*
122. 买卖股票的最佳时机II

给你一个整数数组prices，其中prices[i] 表示某支股票第i天的价格。
在每一天，你可以决定是否购买和/ 或 出售股票。你在任何时候最多只能持有一股 股票。你也可以先购买，然后在同一天出售。
返回你能获得的最大利润

!!! TODO
思路不清
思路：
1. 贪心
2. 动态规划
*/

func maxProfit122(prices []int) int {
	var ans int
	lowPrice := prices[0]
	var curProfit int

	for i := 1; i < len(prices); i++ {
		// 如果当前价格低，卖掉最大利润
		if prices[i] < prices[i-1] || prices[i] < lowPrice {
			if curProfit > 0 {
				ans += curProfit
				curProfit = 0
			}
			lowPrice = prices[i]
			continue
		}
		// 如果价格上涨，更新利润
		if prices[i]-lowPrice > curProfit {
			curProfit = prices[i] - lowPrice
		}
	}
	if curProfit != 0 {
		ans += curProfit
	}
	return ans
}
