package day12

/*
70. 爬楼梯

假设你正在爬楼梯。需要n阶你才能到达楼顶。
每次你可以爬1或2个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定n是一个正整数
*/

func climbStairsOut(n int) int {
	// input 44 超出时间限制
	if n == 1 || n == 2 {
		return n
	}

	return climbStairsOut(n-2) + climbStairsOut(n-1)
}

func climbStairs(n int) int {
	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
