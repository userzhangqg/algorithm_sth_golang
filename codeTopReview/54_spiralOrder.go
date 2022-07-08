package codeTopReview

/*
54. 螺旋矩阵
https://leetcode.cn/problems/spiral-matrix/

复杂度分析

时间复杂度：O(mn)O(mn)，其中 mm 和 nn 分别是输入矩阵的行数和列数。矩阵中的每个元素都要被访问一次。

空间复杂度：O(1)O(1)。除了输出数组以外，空间复杂度是常数。

*/

func spiralOrder(matrix [][]int) []int {
	var ans []int
	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	var x, y int

	for u <= d && l <= r {
		for x, y = u, l; y <= r && u <= d && l <= r; y++ {
			ans = append(ans, matrix[x][y])
		}
		y--
		u++

		for x = x + 1; u <= d && l <= r && x <= d; x++ {
			ans = append(ans, matrix[x][y])
		}
		x--
		r--

		for y = y - 1; u <= d && l <= r && y >= l; y-- {
			ans = append(ans, matrix[x][y])
		}
		y++
		d--

		for x = x - 1; u <= d && l <= r && x >= u; x-- {
			ans = append(ans, matrix[x][y])
		}
		x++
		l++

	}
	return ans
}
