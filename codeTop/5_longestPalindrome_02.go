package codeTop

/*
动态规划解法

时间复杂度O (n^2)
*/

func longestPalindrome2(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start := 0
	size := 1

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			var cur int
			if j-i <= 2 {
				if s[i] == s[j] {
					dp[i][j] = true
					cur = j - i + 1
				}
			} else {
				if s[i] == s[j] && dp[i+1][j-1] {
					dp[i][j] = true
					cur = j - i + 1
				}
			}

			if dp[i][j] && cur > size {
				start = i
				size = cur
			}
		}
	}
	//fmt.Println(dp)
	return s[start : start+size]
}
