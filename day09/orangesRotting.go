package day09

/*
994. 腐烂的橘子

在给定的网格中，每个单元格可以有以下三个值之一
值 0 代表空单元格；
值 1 代表新鲜橘子
值 2 代表腐烂的橘子
每分钟，任何与腐烂的橘子（在4个正方向上）相邻的新鲜橘子都会腐烂。
返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1
*/

/*func orangesRotting(grid [][]int) int {
	res := 0                            // 时间
	queue := []int{}                    // 怀橘子队列
	orange := 0                         // 好橘子个数
	col, row := len(grid[0]), len(grid) // 二维数组的长宽
	// 遍历二维数组,将烂橘子的坐标加入队列
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 统计好橘子数量
			if grid[i][j] == 1 {
				orange++
			}
			// 怀橘子入队
			if grid[i][j] == 2 {
				queue = append(queue, i*col+j)
			}
		}
	}
	// 上下左右四个方向
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	// bfs
	for orange > 0 && len(queue) != 0 {
		size := len(queue) // 提前保存队列长度
		for i := 0; i < size; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			c, r := node%col, node/col
			// 怀橘子的上下左右遍历
			for j := 0; j < 4; j++ {
				nr, nc := r+dx[j], c+dy[j]
				if nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc] == 1 {
					orange--         // 好橘子减一
					grid[nr][nc] = 2 // 好橘子变成坏橘子,然后入队
					queue = append(queue, nr*col+nc)
				}
			}
		}
		// 时间加一
		res++
	}
	if orange != 0 { // 好橘子有剩余
		return -1
	}
	return res
}

*/
type rottCoo struct {
	r int
	c int
	t int
}

func orangesRotting(grid [][]int) int {
	R, C := len(grid)-1, len(grid[0])-1
	fresh, time := 0, 0
	coos := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	//queue := []struct{}
	queue := make([]rottCoo, 0)
	for r, v := range grid {
		for c, k := range v {
			if k == 1 {
				fresh++
			} else if k == 2 {
				queue = append(queue, rottCoo{r, c, time})
			}
		}
	}

	for len(queue) != 0 {
		rot := queue[0]
		queue = queue[1:]
		for _, coo := range coos {
			x := rot.r + coo[0]
			y := rot.c + coo[1]
			if x >= 0 && x <= R && y >= 0 && y <= C && grid[x][y] == 1 {
				grid[x][y] = 2
				fresh--
				queue = append(queue, rottCoo{x, y, rot.t + 1})
				time = rot.t + 1
				//if rot.t + 1 > time {
				//	time = rot.t + 1
				//}
			}
		}
	}

	if fresh != 0 {
		return -1
	}

	return time
}
