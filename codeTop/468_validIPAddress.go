package codeTop

import (
	"strconv"
	"strings"
)

/*
给定一个字符串queryIP。如果是有效的IPv4地址，返回"IPv4"；如果是有效的IPv6地址，返回"IPv6"；如果不是上述类型的IP地址，返回"Neither"。
https://leetcode-cn.com/problems/validate-ip-address/

!!! TODO
思路：
1. 分治判断
*/

func validIPAddress(queryIP string) string {
	v4Num := strings.Count(queryIP, ".")
	v6Num := strings.Count(queryIP, ":")

	if v4Num > 0 && v6Num > 0 {
		return "Neither"
	}
	if v4Num > 0 {
		return checkIPv4(queryIP)
	}
	if v6Num > 0 {
		return checkIPv6(queryIP)
	}
	return "Neither"
}

func checkIPv4(ipStr string) string {
	strArr := strings.Split(ipStr, ".")

	if len(strArr) != 4 {
		return "Neither"
	}

	for i := range strArr {
		if len(strArr[i]) == 0 {
			return "Neither"
		}
		if strArr[i][0] == '0' && len(strArr[i]) > 1 {
			return "Neither"
		}
		curNum, err := strconv.Atoi(strArr[i])
		if err != nil || curNum > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

func checkIPv6(ipStr string) string {
	strArr := strings.Split(ipStr, ":")

	if len(strArr) != 8 {
		return "Neither"
	}

	for i := range strArr {
		if len(strArr[i]) == 0 || len(strArr[i]) > 4 {
			return "Neither"
		}

		for j := range strArr[i] {
			if strArr[i][j] >= 'g' && strArr[i][j] <= 'z' || strArr[i][j] >= 'G' && strArr[i][j] <= 'Z' {
				return "Neither"
			}
		}
	}
	return "IPv6"
}
