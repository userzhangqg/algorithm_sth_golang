package day02

import "sort"

/*
977. 有序数组的平方
给你一个按非递减顺序排序的整数数组nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序
*/

func sortedSquares1(nums []int) []int {
	squs := make([]int, len(nums))
	for i, v := range nums {
		squs[i] = v * v
	}
	sort.Ints(squs)
	return squs
}

func sortedSquares(nums []int) []int {
	/*
		我们可以使用两个指针分别指向位置 00 和 n-1n−1，每次比较两个指针对应的数，选择较大的那个逆序放入答案并移动指针。这种方法无需处理某一指针移动至边界的情况，读者可以仔细思考其精髓所在。
	*/
	squs := make([]int, len(nums))

	i, j := 0, len(nums)-1
	for k := len(nums) - 1; k >= 0; k-- {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			squs[k] = nums[i] * nums[i]
			i++
		} else {
			squs[k] = nums[j] * nums[j]
			j--
		}
	}
	return squs
}
