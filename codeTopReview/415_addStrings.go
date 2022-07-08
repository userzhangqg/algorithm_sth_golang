package codeTopReview

import "strconv"

/*
415. 字符串相加
https://leetcode.cn/problems/add-strings/

复杂度分析

时间复杂度：竖式加法的次数取决于较大数的位数。
空间复杂度：O(1)O(1)。除答案外我们只需要常数空间存放若干变量。在 Java 解法中使用到了 StringBuffer，故 Java 解法的空间复杂度为 O(n)O(n)。
*/

func addStrings(num1 string, num2 string) string {
	var ans string
	var add int

	// 注意判断进位add不等于0
	for i1, i2 := len(num1)-1, len(num2)-1; i1 >= 0 || i2 >= 0 || add != 0; i1, i2 = i1-1, i2-1 {
		var x, y int
		if i1 >= 0 {
			x = int(num1[i1] - '0')
		}
		if i2 >= 0 {
			y = int(num2[i2] - '0')
		}

		sum := x + y + add
		ans = strconv.Itoa(sum%10) + ans
		add = sum / 10
	}
	return ans
}
