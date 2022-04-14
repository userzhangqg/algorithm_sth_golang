package codeTop

func longestCommonPrefixReview(strs []string) string {
	sArr := make([]int, len(strs[0]))
	sLen := len(sArr)

	for _, s := range strs {
		for i := range s {
			if i < sLen && s[i] == strs[0][i] {
				sArr[i]++
			} else {
				break
			}
		}
	}

	var ans string
	for i, v := range sArr {
		if v == len(strs) {
			ans += string(strs[0][i])
		} else {
			break
		}
	}
	return ans
}
