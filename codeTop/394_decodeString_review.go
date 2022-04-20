package codeTop

import "strings"

func decodeStringReview(s string) string {
	var strStack []string
	var numStack []int

	var ans string
	var mul int

	for i := 0; i < len(s); i++ {
		for s[i] >= '0' && s[i] <= '9' {
			mul = mul*10 + int(s[i]-'0')
			i++
		}

		if s[i] == '[' {
			numStack = append(numStack, mul)
			strStack = append(strStack, ans)
			mul = 0
			ans = ""
		} else if s[i] == ']' {
			sTmp := strStack[len(strStack)-1]
			nTmp := numStack[len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
			numStack = numStack[:len(numStack)-1]
			ans = sTmp + strings.Repeat(ans, nTmp)
		} else {
			ans += string(s[i])
		}
	}
	return ans
}
