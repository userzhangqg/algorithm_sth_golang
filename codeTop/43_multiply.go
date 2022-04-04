package codeTop

import "strconv"

/*
43. 字符串相乘

给定两个以字符串形式表示的非负整数nums1和nums2，返回nums1和nums2的乘积，他们的乘积也表示为字符串形式。

注意：不能使用任何内置的BigInteger库或直接将输入转换为整数
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ans string
	len1, len2 := len(num1), len(num2)
	zero := ""
	for i := len2 - 1; i >= 0; i-- {
		add := 0
		curr := ""
		y := int(num2[i] - '0')

		for j := len1 - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			mul := x*y + add
			curr = strconv.Itoa(mul%10) + curr
			add = mul / 10
		}
		if add != 0 {
			curr = strconv.Itoa(add) + curr
		}
		ans = addString(ans, curr+zero)
		zero += "0"
	}
	return ans
}

func addString(str1 string, str2 string) string {
	var add int
	var ans string
	len1, len2 := len(str1), len(str2)
	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
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
