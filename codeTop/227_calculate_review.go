package codeTop

func calculateReview(s string) int {
	var ans int
	var num int
	stack := make([]int, 0)

	preSign := '+'

	for i, v := range s {
		isDigit := v >= '0' && v <= '9'
		if isDigit {
			num = num*10 + int(v-'0')
		}

		if !isDigit && v != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			preSign = v
		}
	}

	for _, v := range stack {
		ans += v
	}
	return ans
}
