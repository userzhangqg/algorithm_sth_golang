package codeTop

import "sort"

func merge56Review(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int
	for _, v := range intervals {
		if len(ans) == 0 || v[0] > ans[len(ans)-1][1] {
			ans = append(ans, v)
		} else if v[1] > ans[len(ans)-1][1] {
			ans[len(ans)-1][1] = v[1]
		}
	}
	return ans
}
