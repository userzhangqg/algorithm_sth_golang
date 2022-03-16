package codeTop

/*
54. 螺旋矩阵
https://leetcode-cn.com/problems/spiral-matrix/

给你一个m 行 n 列的矩阵matrix，请你按照顺时针螺旋顺序，返回矩阵中的所有元素

!!! 思路错误，不可用深度优先
*/

func spiralOrder(matrix [][]int) []int {
	var ans []int
	if len(matrix) == 0 {
		return ans
	}

	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1

	var x, y int
	for u <= d && l <= r {
		for x, y = u, l; y <= r && u <= d && l <= r; y++ {
			ans = append(ans, matrix[x][y])
		}
		u++
		y--

		for x = x + 1; x <= d && u <= d && l <= r; x++ {
			ans = append(ans, matrix[x][y])
		}
		r--
		x--

		for y = y - 1; y >= l && u <= d && l <= r; y-- {
			ans = append(ans, matrix[x][y])
		}
		d--
		y++

		for x = x - 1; x >= u && u <= d && l <= r; x-- {
			ans = append(ans, matrix[x][y])
		}
		l++
		x++
	}
	return ans
}

func dfs54(matrix [][]int, i, j int, ans []int) []int {
	// 错误做法
	if matrix[i][j] == -101 {
		return ans
	}
	ans = append(ans, matrix[i][j])
	matrix[i][j] = -101

	if j < len(matrix[0])-1 {
		ans = dfs54(matrix, i, j+1, ans)
	}
	if i < len(matrix)-1 {
		ans = dfs54(matrix, i+1, j, ans)
	}
	if j > 0 {
		ans = dfs54(matrix, i, j-1, ans)
	}
	if i > 0 {
		ans = dfs54(matrix, i-1, j, ans)
	}
	return ans
}
