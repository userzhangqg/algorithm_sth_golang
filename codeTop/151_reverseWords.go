package codeTop

/*
151. 颠倒字符串中的单词

给你一个字符串s，颠倒字符串中单词的顺序。
单词是由非空格字符串组成的字符串。s中使用至少一个空格将字符串中的单词分隔开。
返回单词顺序颠倒且单词之间用单个空格连接的结果字符串。

注意：输入字符串s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符集中，单词间应当仅使用单个空格分隔，且不包含任何额外的空格。

！！！ 未做出
思路：
https://leetcode-cn.com/problems/reverse-words-in-a-string/solution/151-zong-he-kao-cha-liao-zi-fu-chuan-cao-yr8q/
双反转：
1. 删除多余空格，前后中
2. 翻转整个字符串
3. 翻转单个单词串
*/

func reverseWords(s string) string {
	ss := []byte(s)

	slow, fast := 0, 0

	for len(ss) > 0 && fast < len(ss) && ss[fast] == ' ' {
		fast++
	}

	for ; len(ss) > 0 && fast < len(ss); fast++ {
		if fast > 0 && ss[fast] == ss[fast-1] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}

	if slow > 0 && ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}

	reverse151(ss)
	for i := 0; i < len(ss); {
		// 反转单个单词
		j := i
		// 找到单词的结束位置
		for j < len(ss) && ss[j] != ' ' {
			j++
		}
		// 反转
		reverse151(ss[i:j])
		// 更新下一个单词的起始位置，+1是要跳过单词间的空格
		i = j + 1
	}
	return string(ss)
}

// 为什么不能直接调换?需两数调换
func reverse151(s []byte) {
	n := len(s)

	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

//func reverse151(s []byte){
//	n := len(s)
//	for i:=0;i<n/2;i++{
//		tmp := s[n-1-i]
//		s[n-1-i] = s[i]
//		s[i] = tmp
//	}
//}
