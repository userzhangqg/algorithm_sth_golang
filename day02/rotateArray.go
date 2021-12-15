package day02

import "fmt"

/*
给你一个数组，将数组中的元素向右轮转k个位置，其中 k 是非负数
https://leetcode-cn.com/problems/rotate-array/solution/
*/

func rotate1(nums []int, k int) []int {
	//for i, j := 0, len(nums) - 1; i < j; i, j = i+1, j-1 {
	//	nums[i], nums[j] = nums[j], nums[i]
	//}
	l := len(nums)
	y := k % l
	n := make([]int, l)
	copy(n, nums)
	for i := 0; i < l; i = i + 1 {
		if l-y > i {
			nums[i+y] = n[i]
		} else {
			nums[y-(l-i)] = n[i]
		}
	}
	//fmt.Println(nums)
	return nums
}

func rotate2(nums []int, k int) []int {
	/*
		官方新建数组
	*/
	newNums := make([]int, len(nums))

	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
	fmt.Println(nums)
	return nums
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func rotate(nums []int, k int) []int {
	/*
		翻转再翻转，注意取余和长度为1情况
	*/
	k %= len(nums)
	reverse(nums)
	//fmt.Println(nums)
	reverse(nums[:k])
	//fmt.Println(nums)
	reverse(nums[k:])
	//fmt.Println(nums)
	return nums
}
