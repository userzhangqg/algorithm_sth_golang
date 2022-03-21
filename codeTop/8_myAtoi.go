package codeTop

import (
	"math"
)

/*
8. 字符串转换为整数（atoi）
https://leetcode-cn.com/problems/string-to-integer-atoi/
请你来实现一个myAtoi(string s) 函数，使其能将字符串转换成一个32位有符号整数

注意：
本地中的空白字符只包括字符' '
除前导空格或数字后的其余字符串外，请勿忽略任何其他字符

！！！ 细节错误多，多正负符号，空格，前导零处理
思路：
1. 先处理前导空格
2. 再处理符号
3. 处理数字部分，遇到字符退出
*/
func myAtoi1(s string) int {
	// 双百答案
	sum, sign, i, n := 0, 1, 0, len(s)
	// 丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	// 判定正负
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	// 从左到右依次累加
	for i < n && s[i] >= '0' && s[i] <= '9' {
		sum = 10*sum + int(s[i]-'0')
		// 整数超过32位有符号整数范围,特殊处理
		if sign*sum < math.MinInt32 {
			return math.MinInt32
		} else if sign*sum > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * sum
}

func myAtoi(s string) int {
	var t int
	t = 1
	var ans int
	i, l := 0, len(s)

	for i < l && s[i] == ' ' {
		i++
	}

	// 注意这里用if，只寻找第一个符号
	if i < l {
		if s[i] == '+' {
			t = 1
			i++
		} else if s[i] == '-' {
			t = -1
			i++
		}
	}

	for i < l {
		if s[i] >= '0' && s[i] <= '9' {
			ans = ans*10 + int(s[i]-'0')
			if ans*t > math.MaxInt32 {
				return math.MaxInt32
			}
			if ans*t < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			return ans * t
		}
		i++
	}

	return ans * t
}
