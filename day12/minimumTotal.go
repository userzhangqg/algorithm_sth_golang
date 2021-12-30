package day12

/*
120. 三角形最小路径和

给定一个三角形triangle，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点在这里指的是下标与上一层结点下标相同或者等于上一层结点下标+1的两个结点。
也就是说，如果正位于当前行的下标i，那么下一步可以移动到下一行的下标i或i + 1
*/

func minimumTotal(triangle [][]int) int {
	l := len(triangle)

	var res [][]int
	for i := 1; i <= l; i++ {
		res = append(res, make([]int, i))
	}

	res[0][0] = triangle[0][0]
	for x := 1; x < l; x++ {
		res[x][0] = res[x-1][0] + triangle[x][0]
		for y := 1; y < len(res[x])-1; y++ {
			if res[x-1][y-1] > res[x-1][y] {
				res[x][y] = res[x-1][y] + triangle[x][y]
			} else {
				res[x][y] = res[x-1][y-1] + triangle[x][y]
			}
		}
		res[x][x] = res[x-1][x-1] + triangle[x][x]
	}

	min := res[l-1][0]
	//fmt.Println(res)
	for k := 1; k < len(res[l-1]); k++ {
		if res[l-1][k] < min {
			min = res[l-1][k]
		}
	}
	return min
}
