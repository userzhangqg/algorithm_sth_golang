package codeTop

func minWindowReview(s string, t string) string {
	sMap, tMap := map[byte]int{}, map[byte]int{}

	sLen := len(s)

	for _, v := range t {
		tMap[byte(v)]++
	}

	check := func() bool {
		for k, v := range tMap {
			if sMap[k] < v {
				return false
			}
		}
		return true
	}

	minSize := sLen + 1
	L, R := -1, -1

	for l, r := 0, 0; r < sLen; r++ {
		if tMap[s[r]] > 0 {
			sMap[s[r]]++
		}

		for check() && l <= r {
			// 如果比最小长度小，更新L，R，最小长度
			if r-l+1 < minSize {
				minSize = r - l + 1
				L, R = l, r+1
			}

			// 如果l是公共字符，s缩减一位
			if _, ok := tMap[s[l]]; ok {
				sMap[s[l]]--
			}
			l++
		}
	}
	if L == -1 {
		return ""
	}
	return s[L:R]
}
