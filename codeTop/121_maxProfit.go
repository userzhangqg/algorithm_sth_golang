package codeTop

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

思路：更新最低价格，如果当前价格差最大，更新收益
*/

func maxProfit(prices []int) int {
	ans := 0
	lowPrice := prices[0]
	l := len(prices)
	for i := 1; i < l; i++ {
		if prices[i] < lowPrice {
			lowPrice = prices[i]
			continue
		}
		if prices[i]-lowPrice > ans {
			ans = prices[i] - lowPrice
		}
	}
	return ans
}
