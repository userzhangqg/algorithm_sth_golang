package codeTop

func reverseWordsReview(s string) string {
	ss := []byte(s)
	size := len(ss)
	slow, fast := 0, 0
	for fast < size && ss[fast] == ' ' {
		fast++
	}

	for ; fast < size; fast++ {
		if fast > 0 && ss[fast] == ss[fast-1] && ss[fast] == ' ' {
			continue
		}
		ss[slow] = ss[fast]
		slow++
	}

	// slow已经+1，需判断-1是否是空格
	if ss[slow-1] == ' ' {
		ss = ss[:slow-1]
	} else {
		ss = ss[:slow]
	}

	reverse151Review(ss)

	for i := 0; i < len(ss); i++ {
		j := i
		for j < len(ss) && ss[j] != ' ' {
			j++
		}
		reverse151Review(ss[i:j])
		i = j
	}
	return string(ss)
}

func reverse151Review(ss []byte) {
	size := len(ss)

	for i := 0; i < size/2; i++ {
		ss[i], ss[size-1-i] = ss[size-1-i], ss[i]
	}
}
