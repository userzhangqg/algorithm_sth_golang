package codeTop

import "sort"

/*
56. 合并区间
https://leetcode-cn.com/problems/merge-intervals/
以数组intervals表示若干个区间的集合，其中单个区间为intervals[i] = [start.i, end.i]
请你合并所有重叠的区间，并返回一个 不重叠的区间数组，该数组需恰好覆盖输入中的所有区间


!!! 思路不清晰
思路：
1. 先按左边界从小到大排序，再依次比较右边界
*/

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int

	for _, v := range intervals {
		if len(ans) == 0 || v[0] > ans[len(ans)-1][1] {
			// 没有与前元素重叠，直接添加
			ans = append(ans, v)
		} else if v[1] > ans[len(ans)-1][1] {
			// 重叠后超出结果集右边界，更新结果集右边界
			ans[len(ans)-1][1] = v[1]
		}
	}
	return ans
}
