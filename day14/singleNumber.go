package day14

/*
136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现一次的元素。
*/

func singleNumber(nums []int) int {
	// 相同异或为0，0异或任何数都为其本身
	n := nums[0]

	for i := 1; i < len(nums); i++ {
		n = n ^ nums[i]
	}
	return n
}
