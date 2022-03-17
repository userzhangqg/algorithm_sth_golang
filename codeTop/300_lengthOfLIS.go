package codeTop

/*
300. 最长递增子序列

给你一个整数数组nums，找到其中最长严格递增的子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。


！！！ 思路错误，重点复习
简单替换尾部数据做法错误

思路：
动态规划 + 二分
思想就是让 tail 中存储比较小的元素。这样，tail 未必是真实的最长上升子序列，但长度是对的。
*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)

	for _, v := range nums {
		if len(dp) == 0 || v > dp[len(dp)-1] {
			dp = append(dp, v)
		} else {
			l, r := 0, len(dp)-1
			cur := r
			for l <= r {
				mid := l + (r-l)/2
				if v > dp[mid] {
					l = mid + 1
				} else {
					r = mid - 1
					// 找到小于等于的位置插入
					cur = mid
				}
			}
			dp[cur] = v
		}
	}
	return len(dp)
}

func lengthOfLIS_Error(nums []int) int {
	// 错误思路
	l := len(nums)
	if l == 1 {
		return 1
	}

	end := nums[0]
	ans := 1
	for i := 1; i < l; i++ {
		if nums[i] > end {
			end = nums[i]
			ans++
		} else {
			end = nums[i]
		}
	}
	return ans
}
