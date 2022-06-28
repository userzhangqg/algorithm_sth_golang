package codeTopReview

/*
20. 有效的括号
https://leetcode.cn/problems/valid-parentheses/

思路：栈 + map

复杂度分析

时间复杂度：O(n)O(n)，其中 nn 是字符串 ss 的长度。

空间复杂度：O(n + |\Sigma|)O(n+∣Σ∣)，其中 \SigmaΣ 表示字符集，本题中字符串只包含 66 种括号，|\Sigma| = 6∣Σ∣=6。栈中的字符数量为 O(n)O(n)，而哈希表使用的空间为 O(|\Sigma|)O(∣Σ∣)，相加即可得到总空间复杂度。

*/

func isValid(s string) bool {
	hashMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		size := len(stack)

		w := hashMap[s[i]]

		if size > 0 && stack[size-1] == w {
			stack = stack[:size-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
