package codeTop

/*
69. x的平方根
https://leetcode-cn.com/problems/sqrtx/
给你一个非负整数x，计算并返回x 的算数平方根

由于返回类型是整数，结果只保留整数部分，小数部分将被舍去。
注意：不允许使用任何内置指数函数和算符，例如pow(x, 0.5) 或者x ** 0.5

！！！ 未做出
思路：
二分查找mid ** 2接近x
*/

func mySqrt(x int) int {
	l, r := 0, x

	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
