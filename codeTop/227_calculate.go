package codeTop

/*
227. 基本计算器II

给你一个字符串表达式s，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
你可以假设给定的表达式总是有效的。所有中间结果将在 一定 范围 内。
注意：不允许使用任何字符串作为数学表达式计算的内置函数，必须eval()。

!!! TODO
未做出
思路：
1. 栈实现，转换为加法
*/

func calculate(s string) int {
	stack := []int{}
	preSign := "+"
	num := 0
	var ans int
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}

		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case "+":
				stack = append(stack, num)
			case "-":
				stack = append(stack, -num)
			case "*":
				stack[len(stack)-1] *= num
			case "/":
				stack[len(stack)-1] /= num
			}
			preSign = string(ch)
			num = 0
		}
	}

	for _, v := range stack {
		ans += v
	}
	return ans
}
