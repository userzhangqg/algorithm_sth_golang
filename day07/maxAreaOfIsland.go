package day07

import "fmt"

/*
685. 岛屿的最大面积

给你一个大小为m x n的二进制矩阵grid。
岛屿是由一些相邻的1构成的组合，这里的相邻要求两个1必须在水平或者垂直的四个方向上相邻。你可以假设grid的四个边缘都被0（代表水）包围着。
岛屿的面积是岛上值为1的单元格代数目。
计算并返回grid中最大的岛屿面积。如果没有岛屿，则返回面积为0.
*/

func maxAreaOfIsland(grid [][]int) int {
	maxArea1 := 0
	for i, v := range grid {
		for j, k := range v {
			if k == 1 {
				area := dfs1(grid, i, j, 0)
				fmt.Println(area)
				if area > maxArea1 {
					maxArea1 = area
				}
			}
		}
	}
	return maxArea1
}

func dfs1(grid [][]int, r int, c int, maxArea int) int {
	//fmt.Println(r, c, maxArea)
	if grid[r][c] == 1 {
		maxArea++
		grid[r][c] = 0
	} else {
		return maxArea
	}

	if r > 0 {
		maxArea = dfs1(grid, r-1, c, maxArea)
	}
	if r < len(grid)-1 {
		maxArea = dfs1(grid, r+1, c, maxArea)
	}
	if c > 0 {
		maxArea = dfs1(grid, r, c-1, maxArea)
	}
	if c < len(grid[0])-1 {
		maxArea = dfs1(grid, r, c+1, maxArea)
	}
	return maxArea
}
