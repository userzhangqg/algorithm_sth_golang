package codeTop

/*
70. 爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/

假设你正在爬楼梯。需要n阶你才能到达顶楼。
每次你可以爬1或2个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

func climbStairsError(n int) int {
	// 超时
	if n < 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	x, y := 1, 1

	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
