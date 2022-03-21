package codeTop

import "math"

func myAtoiReview(s string) int {
	i, size := 0, len(s)
	var ans int
	t := 1

	for i < size && s[i] == ' ' {
		i++
	}

	if i < size {
		if s[i] == '-' {
			t = -1
			i++
		} else if s[i] == '+' {
			t = 1
			i++
		}
	}

	for i < size && s[i] >= '0' && s[i] <= '9' {
		ans = ans*10 + int(s[i]-'0')
		if ans*t > math.MaxInt32 {
			return math.MaxInt32
		} else if ans*t < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return ans * t
}
