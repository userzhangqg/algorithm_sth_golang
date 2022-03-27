package codeTop

import (
	"strconv"
	"strings"
)

/*
93. 复原IP地址
有效IP地址正好由四个整数（每个整数位于0 到 255之间组成，且不能含有前导0），整数之间用'.'分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。你 不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-ip-addresses

!!! 无思路
回溯算法总结：https://leetcode-cn.com/problems/restore-ip-addresses/solution/dai-ma-sui-xiang-lu-93-fu-yuan-ip-di-zhi-pzjo/
*/

func restoreIpAddresses(s string) []string {
	var ans, tmp []string
	backtrack93(s, tmp, 0, &ans)
	return ans
}

func backtrack93(s string, tmp []string, start int, ans *[]string) {
	// 终止条件：到结尾，并且ip段正好为四段，添加到ans
	if start == len(s) && len(tmp) == 4 {
		ipString := strings.Join(tmp, ".")
		*ans = append(*ans, ipString)
		return
	}

	for i := start; i < len(s); i++ {
		tmp = append(tmp, s[start:i+1])
		// 如果子段小于等于三位或tmp小于等于四位，递归
		if i+1-start <= 3 && len(tmp) <= 4 && isNormalIp(s, start, i) {
			backtrack93(s, tmp, i+1, ans)
		} else {
			//如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
		// 回溯
		tmp = tmp[:len(tmp)-1]
	}
}

func isNormalIp(s string, start, end int) bool {
	if end-start > 0 && s[start] == '0' {
		// 如果有前导0，剪枝
		return false
	}
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if checkInt > 255 {
		// 如果大于ip限制，剪枝
		return false
	}
	return true
}
