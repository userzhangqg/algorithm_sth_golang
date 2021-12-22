package day09

/*
542. 01 矩阵

给定一个由0和1组成的矩阵mat，请输出一个大小相同的矩阵，其中每一个格子mat中对应位置元素到最近0的距离。
两个相邻元素间的距离为1
*/

func updateMatrix(mat [][]int) [][]int {

	mat1 := make([][]int, len(mat))
	for i, _ := range mat1 {
		mat1[i] = make([]int, len(mat[0]))
		copy(mat1[i], mat[i])
	}

	for i, x := range mat {
		for j, y := range x {
			if y == 0 {
				mat1[i][j] = 0
			} else {
				up := getSmall(mat1, i-1, j)
				left := getSmall(mat1, i, j-1)
				//if up > left {
				//	mat1[i][j] = left
				//}else {
				//	mat1[i][j] = up
				//}
				mat1[i][j] = min(up, left)
			}
			//fmt.Println(i, j, mat1)
		}
	}

	//mat2:=make([][]int,len(mat))
	//for i,_:=range mat2{
	//	mat2[i]=make([]int,len(mat[0]))
	//	copy(mat2[i],mat[i])
	//}
	//mat2 := mat1
	//fmt.Println("")
	for i := len(mat1) - 1; i >= 0; i-- {
		for j := len(mat1[0]) - 1; j >= 0; j-- {
			if mat1[i][j] != 0 {
				down := getSmall(mat1, i+1, j)
				right := getSmall(mat1, i, j+1)
				//if down > right {
				//	mat2[i][j] = right
				//}else {
				//	mat2[i][j] = down
				//}
				mat1[i][j] = min(mat1[i][j], min(down, right))
			}
			//fmt.Println(i, j, mat2)
			//if mat1[i][j] > mat2[i][j] {
			//	mat1[i][j] = mat2[i][j]
			//}
		}
	}

	return mat1

}

func getSmall(mat [][]int, r int, c int) int {
	//if r == -1 {
	//	r = 0
	//}
	//if r == len(mat) {
	//	r = len(mat) - 1
	//}
	//
	//if c == -1 {
	//	c = 0
	//}
	//
	//if c == len(mat[0]) {
	//	c = len(mat[0]) - 1
	//}

	if r == -1 || r == len(mat) || c == -1 || c == len(mat[0]) {
		return 10001
	}

	return mat[r][c] + 1
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
