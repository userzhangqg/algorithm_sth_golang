package day11

import (
	"fmt"
	"unicode"
)

/*
给定一个字符串S， 通过将字符串S中的每个字母转变为大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合
*/

//var res []string

func letterCasePermutation(s string) []string {
	res := findAllString([]rune(s), make([]rune, 0), len(s), make([]string, 0))
	return res
}

func findAllString(s []rune, tmpString []rune, stringLen int, res []string) []string {
	count++
	fmt.Println(count)
	//fmt.Println(string(s), string(tmpString))
	if len(tmpString) == stringLen {
		sliceS := make([]rune, len(tmpString))
		copy(sliceS, tmpString)
		res = append(res, string(sliceS))
		return res
	}

	for i := 0; i < len(s); i++ {
		srcCur := s[i]

		srcString := s
		//srcString := make([]rune, len(s))
		//copy(srcString, s)

		if unicode.IsLetter(srcCur) {
			var cur rune
			if unicode.IsLower(srcCur) {
				cur = unicode.ToUpper(srcCur)
			} else {
				cur = unicode.ToLower(srcCur)
			}

			s = s[i+1:]
			//s = append(s[:i], append([]rune{cur}, s[i+1:]...)...)
			tmpString = append(tmpString, cur)
			res = findAllString(s, tmpString, stringLen, res)
			s = srcString
			//s = append(s[:i], append([]rune{srcCur}, s[i+1:]...)...)
			tmpString = tmpString[:len(tmpString)-1]
		}

		//s = append(s[:i], s[i+1:]...)
		s = s[i+1:]
		tmpString = append(tmpString, srcCur)
		res = findAllString(s, tmpString, stringLen, res)
		s = srcString
		//s = append(s[:i], append([]rune{srcCur}, s[i:]...)...)
		tmpString = tmpString[:len(tmpString)-1]
	}
	return res
}

var count int

func dfs(s []byte, index int, total int, rvl *[]string) {
	count++
	fmt.Println(count)
	if index == total {
		*rvl = append(*rvl, string(s))
		return
	}

	if !(s[index] >= '0' && s[index] <= '9') {
		// lowercase
		s[index] = s[index] | 0x20
		dfs(s, index+1, total, rvl)

		// uppercase
		s[index] = s[index] & 0xDF
		dfs(s, index+1, total, rvl)
	} else {
		dfs(s, index+1, total, rvl)
	}
}
func letterCasePermutationDemo(s string) []string {
	// Demo 为什么调用的次数少？
	raw := []byte(s)
	rvl := []string{}

	dfs(raw, 0, len(raw), &rvl)

	return rvl
}
