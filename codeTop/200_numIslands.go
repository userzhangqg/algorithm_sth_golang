package codeTop

/*
200. 岛屿数量
https://leetcode-cn.com/problems/number-of-islands/

给你一个由1 （陆地）和 0 （水）组成的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。


!!! 需要先遍历遗忘
思路：遍历后递归
*/

func numIslands(grid [][]byte) int {
	ans := 0
	for i, v := range grid {
		for j, k := range v {
			if k == '1' {
				ans++
				dfsNumIslands(grid, i, j)
			}
		}
	}
	return ans
}

func dfsNumIslands(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	} else {
		grid[i][j] = '0'
	}

	if i > 0 {
		dfsNumIslands(grid, i-1, j)
	}
	if i < len(grid)-1 {
		dfsNumIslands(grid, i+1, j)
	}
	if j > 0 {
		dfsNumIslands(grid, i, j-1)
	}
	if j < len(grid[0])-1 {
		dfsNumIslands(grid, i, j+1)
	}

}
