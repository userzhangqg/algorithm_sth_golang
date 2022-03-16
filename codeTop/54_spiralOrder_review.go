package codeTop

func spiralOrderReview(matrix [][]int) []int {
	var ans []int

	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1

	var m, n int
	for u <= d && l <= r {
		m = u
		n = l

		for ; n <= r && u <= d && l <= r; n++ {
			ans = append(ans, matrix[m][n])
		}
		n--
		u++

		for m = m + 1; m <= d && u <= d && l <= r; m++ {
			ans = append(ans, matrix[m][n])
		}
		m--
		r--

		for n = n - 1; n >= l && u <= d && l <= r; n-- {
			ans = append(ans, matrix[m][n])
		}
		n++
		d--

		for m = m - 1; m >= u && u <= d && l <= r; m-- {
			ans = append(ans, matrix[m][n])
		}
		m++
		l++
	}
	return ans
}
