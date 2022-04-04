package codeTop

/*
470. 用Rand7()实现Rand10()
给定方法rand7可生成[1, 7]范围内的均匀随机整数，试写一个方法rand10生成[1, 10]范围内的均匀随机整数。
你只能调用rand7()且不能调用其他方法。请你不要使用系统的Math.random()方法。
每个测试用例将有一个内部参数n，即你实现的函数rand10()在测试时将被调用的次数。请注意，这不是传递给rand10()的参数

!!! TODO

思路：
1. 拒绝采样，舍弃不符合的数据
https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/si-lu-jing-jiang-yong-rand7-shi-xian-ran-9gis/
*/

func rand7() int {
	return 0
}

func rand10() int {
	for {
		m := rand7()
		n := rand7()
		i := (m-1)*7 + n
		if i <= 40 {
			return i%10 + 1
		}
	}
}
