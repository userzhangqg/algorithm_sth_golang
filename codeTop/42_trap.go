package codeTop

/*
42. 接雨水
https://leetcode-cn.com/problems/trapping-rain-water/

给定n个非负整数表示每个宽度为1的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水

！！！ 未做出
https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/

思路：
1. 双指针, 左右寻找短板，用短板-桶深
*/

func trap(height []int) int {
	l, r := 0, len(height)-1
	l_max, r_max := 0, 0
	ans := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= l_max {
				l_max = height[l]
			} else {
				ans = l_max - height[l] + ans
			}
			l++
		} else {
			if height[r] >= r_max {
				r_max = height[r]
			} else {
				ans = r_max - height[r] + ans
			}
			r--
		}
	}
	return ans
}
