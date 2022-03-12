package codeTop

/*
5. 最长回文子串
https://leetcode-cn.com/problems/longest-palindromic-substring/

给你一个字符串s，找到s中最长的回文子串

！！！ 未做出，重点
正确思路：
1. 中心扩散
2. 动态规划
*/

func longestPalindrome(s string) string {
	ans := ""
	for i, n := 0, len(s); i < n; i++ {
		// 枚举每个奇偶长度的回文串的中心起点，找最长的可能性
		l, r := i, i

		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}

		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}
	}
	return ans
}
