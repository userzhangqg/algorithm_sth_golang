package codeTop

/*
136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
你的算法应该具有线性时间复杂度。你可以不适用额外空间来实现吗？

!!! TODO
思路忘记
思路:
1. 位运算，异或运算，相同为0，任何数与0异或为其本身。交换率
*/

func singleNumber(nums []int) int {
	var ans int

	for _, v := range nums {
		ans = ans ^ v
	}
	return ans
}
