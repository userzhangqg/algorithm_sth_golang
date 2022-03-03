package codeTop

/*
3. 无重复字符的最长子串

给定一个字符串s，请你找出其中不含有重复字符的 最长子串 的长度
*/

func lengthOfLongestSubstring(s string) int {
	s_map := map[byte]int{}

	rk, ans := -1, 0

	n := len(s)

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(s_map, s[i-1])
		}

		for rk+1 < n && s_map[s[rk+1]] == 0 {
			s_map[s[rk+1]]++
			rk++
		}

		ans = max(ans, rk+1-i)
	}

	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
