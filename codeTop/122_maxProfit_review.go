package codeTop

func maxProfitReview(prices []int) int {
	var ans int
	// 不限制买卖次数，有利可图就卖掉
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}

	return ans
}
