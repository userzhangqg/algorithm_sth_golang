package day01

/*
704. 二分查找
给定一个n个元素有序的（升序）整形数组nums和一个目标值target，写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回-1
*/

func search(nums []int, target int) int {
	var l, r int = 0, len(nums) - 1

	for l <= r {
		var mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//func main() {
//	nums := []int{-1, 0, 3, 5, 9, 12}
//	target := 9
//	fmt.Println(search(nums, target))
//
//	//nums := []int{-1, 0, 3, 5, 9, 12}
//	//target := 2
//	//fmt.Println(search(nums, target))
//}

//func TestSearch(t *testing.T)  {
//	nums := []int{-1, 0, 3, 5, 9, 12}
//	target := 9
//	if search(nums, target) != 4{
//		t.Errorf("error")
//	}
//}
