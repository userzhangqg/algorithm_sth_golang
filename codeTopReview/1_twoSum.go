package codeTopReview

/*
1. 两数之和
https://leetcode.cn/problems/two-sum/

未作出：haspMap
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
