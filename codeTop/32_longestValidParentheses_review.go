package codeTop

func longestValidParenthesesReview(s string) int {
	var ans int
	var q []int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || len(q) == 0 {
			q = append(q, i)
		} else {
			x := q[len(q)-1]
			if s[x] == '(' {
				q = q[:len(q)-1]
				if len(q) == 0 {
					ans = i + 1
				} else {
					ans = max32Review(ans, i-q[len(q)-1])
				}
			} else {
				q = append(q, i)
			}
		}
	}
	return ans
}

func max32Review(x, y int) int {
	if x > y {
		return x
	}
	return y
}
