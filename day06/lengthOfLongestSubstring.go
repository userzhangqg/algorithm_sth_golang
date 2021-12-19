package day06

/*
3. 无重复字符的最长子串

给定一个字符串s，请你找出其中不含有重复字符的 最长子串 的长度
*/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	l, r, k := 0, len(s)-1, 0
	longest := 0

	//if r == 0 {
	//	return 1
	//}

	for l <= r {
		if l > 0 {
			delete(m, s[l-1])
		}

		for k <= r && m[s[k]] == 0 {
			m[s[k]]++
			k++
		}

		// 子串长度应该由右指针 - 左指针
		longest = max(longest, k-l)

		l++
	}

	return longest
}

func lengthOfLongestSubstringDemo(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
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
