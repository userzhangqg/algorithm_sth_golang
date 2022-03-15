package codeTop

import "strconv"

/*
415. 字符串相加
https://leetcode-cn.com/problems/add-strings/

给定两个字符串形式的非负整数nums1 和 nums2， 计算它们的和并同样以字符串形式返回。
你不能使用任何内建的用于处理大整数的库，也不能直接将输入的字符串转换为整数形式。

!!! 未做出
思路：竖式加法
*/

func addStrings(num1 string, num2 string) string {
	ans := ""
	add := 0

	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + add
		ans = strconv.Itoa(sum%10) + ans
		add = sum / 10
	}
	return ans
}
