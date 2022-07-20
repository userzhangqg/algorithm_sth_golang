package codeTopReview

import "sort"

/*
56. 合并区间
https://leetcode.cn/problems/merge-intervals/

复杂度分析

时间复杂度：
O(nlogn)，其中
n 为区间的数量。除去排序的开销，我们只需要一次线性扫描，所以主要的时间开销是排序的
O(nlogn)。
空间复杂度：
O(logn)，其中
n 为区间的数量。这里计算的是存储答案之外，使用的额外空间。
O(logn) 即为排序所需要的空间复杂度。
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
