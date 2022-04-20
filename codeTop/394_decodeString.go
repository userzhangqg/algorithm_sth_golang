package codeTop

import (
	"strings"
)

/*
394. 字符串解码
https://leetcode-cn.com/problems/decode-string/
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。

!!! TODO
思路错误：
思路：
1. 双栈
https://leetcode-cn.com/problems/decode-string/solution/zhan-de-ji-yi-nei-ceng-de-jie-ma-liao-bie-wang-lia/
*/

func decodeString(s string) string {
	strStack := make([]string, 0)
	numStack := make([]int, 0)

	ans := ""

	for i := 0; i < len(s); i++ {
		var num int
		for s[i] >= '0' && s[i] <= '9' {
			n := s[i] - '0'
			num = num*10 + int(n)
			i++
		}
		if s[i] == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, ans)
			ans = ""
		} else if s[i] == ']' {
			sTmp := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			nTmp := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			ans = sTmp + strings.Repeat(ans, nTmp)
		} else {
			ans += string(s[i])
		}
	}
	return ans
}

/*
func decodeString(s string) string {
	var ans string
	var tmp string
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ']' {
			count++
		}else if s[i] == '[' {
			//mul := s[i-1] - '0'
			//i--
			var mul string
			for i > 0 && s[i-1] - '0' <= 10 {
				mul = string(s[i-1]) + mul
				i--
			}
			//i++
			m, _ := strconv.Atoi(mul)
			tmp = strings.Repeat(tmp, m)
			count--
			if count == 0 {
				ans = tmp + ans
				tmp = ""
			}
		}else if count == 0 {
			ans = string(s[i]) + ans
		}else {
			tmp = string(s[i]) + tmp
		}
	}
	return ans
}

*/
