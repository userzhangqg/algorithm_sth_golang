package codeTop

func quickSort2(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort2(sortArray, left, i-1)
		quickSort2(sortArray, j+1, right)
	}
}

//func main(){
//	nums := []int{3,2,3,1,2,4,5,5,6}
//	quickSort2(nums, 0, len(nums) - 1)
//	fmt.Println(nums)
//}
