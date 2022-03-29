package codeTop

/*
41. 缺失的第一个正数
https://leetcode-cn.com/problems/first-missing-positive/

给你一个未排序的整数数组nums，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为O(n)并且只使用常数级别额外空间的解决方案。

！！！ 未做出
将数组设计成哈希表的思路：

我们对数组进行遍历，对于遍历到的数 xx，如果它在 [1, N][1,N] 的范围内，那么就将数组中的第 x-1x−1 个位置（注意：数组下标从 00 开始）打上「标记」。在遍历结束之后，如果所有的位置都被打上了标记，那么答案是 N+1N+1，否则答案是最小的没有打上标记的位置加 11。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution/
*/

func firstMissingPositive(nums []int) int {
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			nums[i] = size + 1
		}
	}

	for i := 0; i < size; i++ {
		// 可能符号被变掉，使用绝对值
		n := abs41(nums[i])
		if n <= size {
			nums[n-1] = -abs41(nums[n-1])
		}
	}

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return size + 1
}

func abs41(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
