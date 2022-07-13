package codeTopReview

/*
42. 接雨水
https://leetcode.cn/problems/trapping-rain-water/

复杂度分析:

时间复杂度：O(n)O(n)，其中 nn 是数组 \textit{height}height 的长度。两个指针的移动总次数不超过 nn。
空间复杂度：O(1)O(1)。只需要使用常数的额外空间。

*/

func trap(height []int) int {
	l, r := 0, len(height)-1
	var l_max, r_max, ans int

	for l < r {
		if height[l] < height[r] {
			if height[l] >= l_max {
				l_max = height[l]
			} else {
				ans = ans + l_max - height[l]
			}
			l++
		} else {
			if height[r] >= r_max {
				r_max = height[r]
			} else {
				ans = ans + r_max - height[r]
			}
			r--
		}
	}
	return ans
}
