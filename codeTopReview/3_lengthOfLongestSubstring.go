package codeTopReview

/*
3. 无重复字符的最长子串
https://leetcode.cn/problems/longest-substring-without-repeating-characters/

Solution:
1. 滑动窗口，map实现
*/

func lengthOfLongestSubstring(s string) int {
	s_map := map[byte]int{}

	c_index, ans := -1, 0

	n := len(s)

	for i := 0; i < n; i++ {
		if i > 0 {
			delete(s_map, s[i-1])
		}

		for c_index+1 < n && s_map[s[c_index+1]] == 0 {
			s_map[s[c_index+1]]++
			c_index++
		}

		ans = max(ans, c_index-i+1)
	}
	return ans
}
