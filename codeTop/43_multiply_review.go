package codeTop

import "strconv"

func multiplyReview(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ans string
	var zero string
	len1, len2 := len(num1), len(num2)
	for i := len2 - 1; i >= 0; i-- {
		var tmp string
		var add int
		x := int(num2[i] - '0')

		for j := len1 - 1; j >= 0; j-- {
			y := int(num1[j] - '0')
			sum := x*y + add
			tmp = strconv.Itoa(sum%10) + tmp
			add = sum / 10
		}
		if add != 0 {
			tmp = strconv.Itoa(add) + tmp
		}
		ans = addString43Review(ans, tmp+zero)
		zero += "0"
	}
	return ans
}

func addString43Review(str1 string, str2 string) string {
	var ans string
	var add int

	for i, j := len(str1)-1, len(str2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(str1[i] - '0')
		}
		if j >= 0 {
			y = int(str2[j] - '0')
		}
		sum := x + y + add
		ans = strconv.Itoa(sum%10) + ans
		add = sum / 10
	}
	return ans
}
