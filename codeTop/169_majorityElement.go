package codeTop

/*
169. 多数元素
给定一个大小为n的数组，找到其中的多数元素。多数元素是值在数组中出现的次数大于 n/2 的元素
你可以假设数组是非空的，并且给定的数组总是存在多数元素

!!! TODO
思路：
1. 摩尔投票法，这方法只有在众数过半的时候才能用，不然很可能选不出最多的那个数。
算法可以分为两个阶段：
对抗阶段：分属两个候选人的票数进行两两对抗抵消
计数阶段：计算对抗结果中最后留下的候选人票数是否有效
https://leetcode-cn.com/problems/majority-element/solution/tu-jie-mo-er-tou-piao-fa-python-go-by-jalan/
*/

func majorityElement(nums []int) int {
	var major int
	var count int

	for _, v := range nums {
		if count == 0 {
			major = v
		}

		if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}
