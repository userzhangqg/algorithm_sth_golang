package codeTop

func isValidReview(s string) bool {
	hashMap := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var q []string

	for _, i := range s {
		if len(q) > 0 && hashMap[string(i)] == q[len(q)-1] {
			q = q[:len(q)-1]
		} else {
			q = append(q, string(i))
		}
	}
	if len(q) == 0 {
		return true
	} else {
		return false
	}
}
