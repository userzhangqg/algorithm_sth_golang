/*
已作答：1,2,4

正确答案：0,1,3

考点：数组长度增加，数组引用
切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，
http://c.biancheng.net/view/28.html

引用传递，值传递
https://blog.csdn.net/baolingye/article/details/111142386
数组默认值传递，使用引用传递用*
切片使用引用传递

数组，切片区别：
数组长度固定

数组和slice的区别
声明数组时，方括号内写明了数组的长度或者...,声明slice时候，方括号内为空
作为函数参数时，数组传递的是数组的副本，而slice传递的是指针
*/

package main

import "fmt"

func TestSlice(arr []int) {
	arr = append(arr, 3)
	arr = append(arr, 4)
}

func main() {
	arr := make([]int, 1, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	TestSlice(arr)
	fmt.Printf("%d,%d,%d", arr[0], arr[1], len(arr))
}
