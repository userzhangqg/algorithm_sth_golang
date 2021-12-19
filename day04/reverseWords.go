package day04

import "strings"

func reverseWords(s string) string {
	var strArray []string
	for _, v := range strings.Split(s, " ") {
		w := reverseS([]byte(v))
		strArray = append(strArray, w)
	}
	return strings.Join(strArray, " ")
}

func reverseS(s []byte) string {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
	return string(s)
}
