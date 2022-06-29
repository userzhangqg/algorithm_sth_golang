package codeTopReview

/*
121. 买卖股票的最佳时机
https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

复杂度分析

时间复杂度：O(n)O(n)，只需要遍历一次。
空间复杂度：O(1)O(1)，只使用了常数个变量。
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

//func maxProfit(prices []int) int {
//	lowPrice := prices[0]
//	prices[0] = 0
//
//	for i := 1; i < len(prices); i++ {
//		if prices[i] < lowPrice {
//			lowPrice = prices[i]
//		}
//		if prices[i] - lowPrice > prices[i-1] {
//			prices[i] = prices[i] - lowPrice
//		}else {
//			prices[i] = prices[i-1]
//		}
//	}
//	return prices[len(prices)-1]
//}
