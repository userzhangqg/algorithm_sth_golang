package codeTop

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

思路：map + 栈
*/

func isValid(s string) bool {
	hashMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	q := []string{}
	for _, i := range s {
		w := hashMap[string(i)]

		size := len(q)
		if size > 0 && w == q[size-1] {
			q = q[:size-1]
		} else {
			q = append(q, string(i))
		}
	}

	if len(q) == 0 {
		return true
	} else {
		return false
	}
}
