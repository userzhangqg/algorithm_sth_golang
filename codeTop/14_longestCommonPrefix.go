package codeTop

/*
14. 最长公共前缀

编写一个函数来查找字符串数组中最长公共前缀
如果不存在公共前缀，返回空字符串""

$ Done
思路：
1. 横向扫描
*/

func longestCommonPrefix(strs []string) string {
	// 最长公共前缀不会长于任意一个字符的长度
	sArr := make([]int, len(strs[0]))
	var ans string

	for _, s := range strs {
		for j := range s {
			if j < len(sArr) && strs[0][j] == s[j] {
				sArr[j]++
			} else {
				break
			}
		}
	}

	for i, v := range sArr {
		if v == len(strs) {
			ans += string(strs[0][i])
		}
	}
	return ans
}
