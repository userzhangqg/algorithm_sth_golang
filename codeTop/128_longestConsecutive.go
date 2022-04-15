package codeTop

/*
128. 最长连续序列
https://leetcode-cn.com/problems/longest-consecutive-sequence/
给定一个未排序的整数数组nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为O(n)的算法并解决此问题。

!!! TODO
思路：
1. hashMap
*/

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)

	for _, num := range nums {
		hashMap[num] = true
	}

	var ans int

	for i := range hashMap {
		// 如果该数是连续数最小起点，从该起点开始增加判断
		// 外层循环需要 O(n)O(n) 的时间复杂度，只有当一个数是连续序列的第一个数的情况下才会进入内层循环，然后在内层循环中匹配连续序列中的数，因此数组中的每个数只会进入内层循环一次。根据上述分析可知，总时间复杂度为 O(n)O(n)，符合题目要求。
		if !hashMap[i-1] {
			curLong := 1
			curIndex := i
			for hashMap[curIndex+1] {
				curLong++
				curIndex++
			}

			if curLong > ans {
				ans = curLong
			}
		}
	}
	return ans
}
