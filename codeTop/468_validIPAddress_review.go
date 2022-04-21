package codeTop

import (
	"strconv"
	"strings"
)

func validIPAddressReview(queryIP string) string {
	countV4 := strings.Count(queryIP, ".")
	countV6 := strings.Count(queryIP, ":")
	if countV4 > 0 && countV6 > 0 {
		return "Neither"
	}
	if countV4 > 0 {
		return validIPv4(queryIP)
	}
	if countV6 > 0 {
		return validIPv6(queryIP)
	}
	return "Neither"
}

func validIPv4(IPStr string) string {
	IPArr := strings.Split(IPStr, ".")

	if len(IPArr) != 4 {
		return "Neither"
	}
	for i := range IPArr {
		if IPArr[i] == "" {
			return "Neither"
		}
		// 此处review错误， > 0 应为 > 1
		if IPArr[i][0] == '0' && len(IPArr[i]) > 1 {
			return "Neither"
		}
		n, err := strconv.Atoi(IPArr[i])
		if err != nil || n > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

func validIPv6(IPStr string) string {
	IPArr := strings.Split(IPStr, ":")

	if len(IPArr) != 8 {
		return "Neither"
	}
	for i := range IPArr {
		// 此处错误，还需限定长度不超过4
		if IPArr[i] == "" || len(IPArr[i]) > 4 {
			return "Neither"
		}
		for j := range IPArr[i] {
			if IPArr[i][j] >= 'g' && IPArr[i][j] <= 'z' || IPArr[i][j] >= 'G' && IPArr[i][j] <= 'Z' {
				return "Neither"
			}
		}
	}
	return "IPv6"
}
