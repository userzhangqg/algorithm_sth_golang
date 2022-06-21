package codeTop

func numIslandsReview(grid [][]byte) int {
	ans := 0
	for i, v := range grid {
		for j, k := range v {
			if k == '1' {
				ans++
			}
			dfs200(grid, i, j)
		}
	}
	return ans
}

func dfs200(grid [][]byte, i, j int) {
	if grid[i][j] == '0' {
		return
	} else {
		grid[i][j] = '0'
	}

	if i > 0 {
		dfs200(grid, i-1, j)
	}
	if i < len(grid)-1 {
		dfs200(grid, i+1, j)
	}
	if j > 0 {
		dfs200(grid, i, j-1)
	}
	if j < len(grid[0])-1 {
		dfs200(grid, i, j+1)
	}
}
