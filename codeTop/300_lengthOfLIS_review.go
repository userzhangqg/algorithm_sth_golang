package codeTop

/*
待二轮复习，边界不清晰
*/

func lengthOfLISReview(nums []int) int {
	var dp []int

	for _, v := range nums {
		if len(dp) == 0 || v > dp[len(dp)-1] {
			dp = append(dp, v)
			continue
		}
		l, r, cur := 0, len(dp)-1, len(dp)-1

		// 注意应该<=
		for l <= r {
			mid := l + (r-l)/2
			if v > dp[mid] {
				l = mid + 1
			} else {
				// 找到不大于原序列元素的位置
				cur = mid
				r = mid - 1
			}
		}
		dp[cur] = v
	}
	return len(dp)
}
