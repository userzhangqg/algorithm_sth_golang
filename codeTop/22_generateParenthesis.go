package codeTop

import "strings"

/*
22. 括号生成

数字n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的符号组合
*/

func generateParenthesis(n int) []string {
	var tmp, ans []string
	backtrack22(0, 0, n, tmp, &ans)
	return ans
}

func backtrack22(i int, j int, n int, tmp []string, ans *[]string) {
	if len(tmp) == n*2 {
		s := strings.Join(tmp, "")
		*ans = append(*ans, s)
		return
	}

	if i < n {
		backtrack22(i+1, j, n, append(tmp, "("), ans)
	}

	if j < i {
		backtrack22(i, j+1, n, append(tmp, ")"), ans)
	}
}
