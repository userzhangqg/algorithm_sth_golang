package codeTopReview

/*
200. 岛屿数量
https://leetcode.cn/problems/number-of-islands/

复杂度分析

时间复杂度：O(MN)O(MN)，其中 MM 和 NN 分别为行数和列数。

空间复杂度：O(MN)O(MN)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 M NMN。

*/

func numIslands(grid [][]byte) int {
	var num int

	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if grid[i][j] == '1' {
			grid[i][j] = '0'
		} else {
			return
		}

		if i > 0 {
			dfs(grid, i-1, j)
		}
		if i < len(grid)-1 {
			dfs(grid, i+1, j)
		}
		if j > 0 {
			dfs(grid, i, j-1)
		}
		if j < len(grid[0])-1 {
			dfs(grid, i, j+1)
		}
	}

	for i, v := range grid {
		for j, k := range v {
			if k == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}

	return num
}
