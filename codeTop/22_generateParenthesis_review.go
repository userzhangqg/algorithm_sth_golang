package codeTop

import "strings"

func generateParenthesisReview(n int) []string {
	var res, tmp []string
	backtrack22Review(0, 0, n, tmp, &res)
	return res
}

func backtrack22Review(i int, j int, n int, tmp []string, res *[]string) {
	if len(tmp) == n*2 {
		s := strings.Join(tmp, "")
		*res = append(*res, s)
		return
	}

	if i < n {
		backtrack22Review(i+1, j, n, append(tmp, "("), res)
	}
	if j < i {
		backtrack22Review(i, j+1, n, append(tmp, ")"), res)
	}
}
