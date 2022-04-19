package codeTop

/*
695. 岛屿的最大面积

给你一个大小为m x n 的二进制矩阵grid
岛屿 是由一些相邻的1构成的组合，你可以假设grid的四个边缘都被0 包围着。
岛屿的最大面积是岛上为1的单元格的数目。
计算并返回grid中最大的岛屿面积。如果没有岛屿，则返回面积为0

$ Done
*/

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				var cur int
				dfs695(grid, &cur, &ans, i, j)
			}
		}
	}
	return ans
}

func dfs695(grid [][]int, curArea *int, maxArea *int, r int, c int) {
	if grid[r][c] == 1 {
		*curArea++
		grid[r][c] = 0
	} else {
		return
	}

	if r < len(grid)-1 {
		dfs695(grid, curArea, maxArea, r+1, c)
	}
	if r > 0 {
		dfs695(grid, curArea, maxArea, r-1, c)
	}
	if c < len(grid[0])-1 {
		dfs695(grid, curArea, maxArea, r, c+1)
	}
	if c > 0 {
		dfs695(grid, curArea, maxArea, r, c-1)
	}

	if *curArea > *maxArea {
		*maxArea = *curArea
	}
}
