package codeTopReview

/*
70. 爬楼梯
https://leetcode.cn/problems/climbing-stairs/

复杂度分析

时间复杂度：循环执行n 次，每次花费常数的时间代价，故渐进时间复杂度为O(n)。
空间复杂度：这里只用了常数个变量作为辅助空间，故渐进空间复杂度为O(1)。

*/

func climbStairs(n int) int {
	x, y := 1, 1

	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
