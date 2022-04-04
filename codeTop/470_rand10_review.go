package codeTop

func rand10Review() int {
	for {
		r := rand7()
		c := rand7()
		res := (r-1)*7 + c
		if res <= 40 {
			return res%10 + 1
		}
	}
}
