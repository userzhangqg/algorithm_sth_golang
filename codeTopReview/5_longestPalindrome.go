package codeTopReview

/*
5. 最长回文子串
https://leetcode.cn/problems/longest-palindromic-substring/

思路：
中心扩散，注意奇偶

复杂度分析

时间复杂度：O(n^2) ，其中 nn 是字符串的长度。长度为 11 和 22 的回文中心分别有 nn 和 n-1n−1 个，每个回文中心最多会向外扩展 O(n)O(n) 次。

空间复杂度：O(1)O(1)。

*/

func longestPalindrome(s string) string {
	ans := ""

	for i, n := 0, len(s); i < n; i++ {
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
