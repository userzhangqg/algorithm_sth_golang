package codeTop

/*
239. 滑动窗口最大值

给你一个整数数组nums，有一个大小为k的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的k个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

！！！ 未做出
思路：
https://leetcode-cn.com/problems/sliding-window-maximum/solution/dong-hua-yan-shi-dan-diao-dui-lie-239hua-hc5u/
1. 双端队列
*/

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	size := len(nums)
	if k < 0 || k > size || size == 0 {
		return ans
	}

	var q []int

	for right := 0; right < size; right++ {
		for len(q) > 0 && nums[right] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, right)

		left := right - k + 1
		if q[0] < left {
			q = q[1:]
		}

		if right+1 >= k {
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}
