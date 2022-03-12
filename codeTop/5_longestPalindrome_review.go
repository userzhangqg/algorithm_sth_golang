package codeTop

func longestPalindromeReview(s string) string {
	var ans string

	for i, n := 0, len(s); i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}

		if r-(l+1) > len(ans) {
			ans = s[l+1 : r]
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}

		if r-(l+1) > len(ans) {
			ans = s[l+1 : r]
		}
	}
	return ans
}
