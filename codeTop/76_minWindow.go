package codeTop

/*
76. 最小覆盖子串
https://leetcode-cn.com/problems/minimum-window-substring/

给你一个字符串s、一个字符串t。返回s中涵盖t所有字符的最小子串。
如果s中不存在涵盖t所有字符的子串，则返回空字符串""。

注意：
- 对于t中重复字符，我们寻找的字符串中该字符数量必须不少于t中该字符的数量
- 如果s中存在这样的子串，我们保证他是唯一的答案

*/

func minWindow(s string, t string) string {
	sMap, tMap := map[byte]int{}, map[byte]int{}
	sLen := len(s)

	for _, v := range t {
		tMap[byte(v)]++
	}

	size := sLen + 1
	L, R := -1, -1

	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}

		for check() && l <= r {
			if r-l+1 < size {
				size = r - l + 1
				L, R = l, l+size
			}

			if _, ok := tMap[s[l]]; ok {
				sMap[s[l]]--
			}
			l++
		}
	}

	if L == -1 {
		return ""
	}
	return s[L:R]
}
