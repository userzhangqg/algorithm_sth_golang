package codeTop

func maxAreaOfIslandReview(grid [][]int) int {
	var ans int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				var curArea int
				dfs695Review(grid, &curArea, &ans, i, j)
			}
		}
	}
	return ans
}

func dfs695Review(grid [][]int, curArea *int, maxArea *int, r int, c int) {
	if grid[r][c] == 1 {
		grid[r][c] = 0
		*curArea++
	} else {
		return
	}

	if r < len(grid)-1 {
		dfs695Review(grid, curArea, maxArea, r+1, c)
	}
	if r > 0 {
		dfs695Review(grid, curArea, maxArea, r-1, c)
	}
	if c < len(grid[0])-1 {
		dfs695Review(grid, curArea, maxArea, r, c+1)
	}
	if c > 0 {
		dfs695Review(grid, curArea, maxArea, r, c-1)
	}

	if *curArea > *maxArea {
		*maxArea = *curArea
	}
}
