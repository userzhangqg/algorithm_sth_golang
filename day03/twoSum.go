package day03

/*
167. 两数之和II-输入有序数组
给定一个已按照 非递减顺序排列 的整数数组 numbers，请你从数组中找出两个数满足相加之和等于目标数target。
函数应该以长度为2的数组数组的形式返回这两个数的下标值。 numbers的下标从1开始计数
你可以假设每个输入 只对应唯一的答案， 而且你不可以重复使用相同的元素。
*/

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return make([]int, 2)
}
