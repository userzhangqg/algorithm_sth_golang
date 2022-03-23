package main

import (
	"fmt"
	"unsafe"
)

/*
已作答：0

正确答案：24

知识点：
https://blog.csdn.net/HaoDaWang/article/details/80005072
如果x为一个切片，sizeof返回的大小是切片的描述符，而不是切片所指向的内存的大小。上面声明了一个切片，然后打印出sizeof的值为24，但是不管slice里的元素为多少，sizeof返回的数据都是24。
*/

type Item struct {
	books []string
}

func main() {
	d := Item{}
	fmt.Printf("%d", unsafe.Sizeof(d.books))
}
