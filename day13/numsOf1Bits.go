package day13

/*
191. 位1的个数

编写一个函数，输入是一个无符号的整数，以二进制串的形式
返回其二进制表达式中位数为 1 的个数（也称为汉明重量）
*/

func hammingWeight(num uint32) int {
	// 任何数和0 做与运算为 0
	var sum int
	for i := 0; i < 32; i++ {
		if num&(1<<i) != 0 {
			sum++
		}
	}
	return sum
}
