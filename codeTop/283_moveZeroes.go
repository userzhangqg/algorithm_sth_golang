package codeTop

/*
283. 移动零

给定一个数组nums，编写一个函数将所有的0移动到数组的末尾，同时保持非零元素的相对顺序
请注意，必须在不复制数组的情况下原地对数组进行操作

!!! TODO
做过思路忘记
思路：
1. 快慢指针，非零时调换
*/

func moveZeroes(nums []int) {
	size := len(nums)
	for left, right := 0, 0; right < size; right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
