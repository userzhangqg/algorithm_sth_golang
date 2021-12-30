package day13

/*
231. 2的幂

给你一个整数n，请你判断该整数是否是2的幂次方。如果是，返回true；否则，返回false。
如果存在一个整数x使得n == 2^x, 则认为n 是 2 的幂次方。
*/

func isPowerOfTwo(n int) bool {
	// 二进制进位性质，如果为幂次方，减1后那么低位全部为1，与运算后为0
	if n > 0 && n&(n-1) == 0 {
		return true
	} else {
		return false
	}
}
