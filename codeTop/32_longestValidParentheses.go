package codeTop

/*
32. 最长有效括号

给你一个只包含'('和')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度

!!! 未做出
注意：为连续有效括号
思路：
1. 栈
*/

func longestValidParentheses(s string) int {
	var stack []int
	var maxLen int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || len(stack) == 0 {
			// 栈为空入栈
			stack = append(stack, i)
		} else {
			// 栈顶元素为左括号，出栈
			if s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					// 空栈的话需计算全长
					maxLen = i + 1
				} else {
					maxLen = max32(maxLen, i-stack[len(stack)-1])
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return maxLen
}

func max32(x, y int) int {
	if x > y {
		return x
	}
	return y
}
