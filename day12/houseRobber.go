package day12

/*
198. 打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统
如果两间相邻的房屋在同一个晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下，一夜之内能够偷窃到的最高金额。
*/

func robOut(nums []int) int {
	// 超出时间限制
	l := len(nums)

	if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	} else {
		a := robOut(nums[:l-1])
		b := nums[l-1] + robOut(nums[:l-2])
		if a > b {
			return a
		} else {
			return b
		}
	}
}

func rob(nums []int) int {

	l := len(nums)

	if l == 1 {
		return nums[0]
	} else {
		if nums[0] > nums[1] {
			nums[1] = nums[0]
		}
	}

	for i := 2; i < l; i++ {
		if nums[i-1] > nums[i]+nums[i-2] {
			nums[i] = nums[i-1]
		} else {
			nums[i] = nums[i] + nums[i-2]
		}
	}

	return nums[l-1]
}
