package day14

/*
190. 颠倒二进制位

颠倒给定的32位无符号整数的二进制位
*/

func reverseBits(num uint32) uint32 {
	// num 与 1 与取最后一位，然后右移；res 左移 与 num移出位进行或运算
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num = num >> 1
	}

	return res
}
