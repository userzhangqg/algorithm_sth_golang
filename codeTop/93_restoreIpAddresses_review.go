package codeTop

import (
	"strconv"
	"strings"
)

func restoreIpAddressesReview(s string) []string {
	var ans, tmp []string
	backtrack93Review(s, tmp, 0, &ans)
	return ans
}

func backtrack93Review(s string, tmp []string, start int, ans *[]string) {
	if start == len(s) && len(tmp) == 4 {
		ipString := strings.Join(tmp, ".")
		*ans = append(*ans, ipString)
		return
	}

	for i := start; i < len(s); i++ {
		tmp = append(tmp, s[start:i+1])
		if i-start+1 <= 3 && len(tmp) <= 4 && isNormalIpReview(s, start, i) {
			backtrack93Review(s, tmp, i+1, ans)
		} else {
			return
		}
		tmp = tmp[:len(tmp)-1]
	}
}

func isNormalIpReview(s string, start, end int) bool {
	if end-start > 0 && s[start] == '0' {
		return false
	}
	subIp, _ := strconv.Atoi(s[start : end+1])
	if subIp > 255 {
		return false
	}
	return true
}
