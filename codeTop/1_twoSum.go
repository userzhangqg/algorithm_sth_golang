package codeTop

/*
1. 两数之和
https://leetcode-cn.com/problems/two-sum/

给定一个整数数组nums和一个整数目标值target，请你在该数组中找出 和为目标值 target 的那 两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案

!!!未做出

正确思路：哈希表

*/

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, v := range nums {
		if k, ok := hashMap[target-v]; ok {
			return []int{k, i}
		}
		hashMap[v] = i
	}
	return nil
}
