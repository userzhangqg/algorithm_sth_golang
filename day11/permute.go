package day11

/*
46. 全排列

给定一个含重复数字的数组nums，返回其所有可能的全排列。你可以按任意顺序返回答案
*/

func permute(nums []int) [][]int {
	var res [][]int
	res = findAll(nums, make([]int, 0), res)
	return res
}

func findAll(nums []int, tmpNums []int, res [][]int) [][]int {
	if len(nums) == 0 {
		slice := make([]int, len(tmpNums))
		copy(slice, tmpNums)
		res = append(res, slice)
		return res
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		tmpNums = append(tmpNums, cur)
		nums = append(nums[:i], nums[i+1:]...)
		res = findAll(nums, tmpNums, res)
		// 注意这里从i开始，因为nums已经变化
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		tmpNums = tmpNums[:len(tmpNums)-1]
	}

	return res
}
