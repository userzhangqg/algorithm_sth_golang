package day06

import "reflect"

/*
567. 字符串的排列

给你两个字符串s1和s2，写一个函数来判断s2是否包含s1的排列。
如果是，返回true；否则，返回false。
换句话说，s1的排列之一是s2的子串。
*/

func checkInclusion(s1 string, s2 string) bool {
	/*
		判断两个子串相同元素的个数及数量是否相同。
	*/
	n, m := len(s1), len(s2)

	if n > m {
		return false
	}

	cnt1, cnt2 := make([]int, 26), make([]int, 26)
	//var cnt1 [26]int
	//var cnt2 [26]int

	for i, v := range s1 {
		cnt1[v-'a']++
		cnt2[s2[i]-'a']++
	}

	if reflect.DeepEqual(cnt1, cnt2) {
		return true
	}

	for j := n; j < m; j++ {
		cnt2[s2[j]-'a']++
		cnt2[s2[j-n]-'a']--

		if reflect.DeepEqual(cnt1, cnt2) {
			return true
		}
	}

	return false

}
