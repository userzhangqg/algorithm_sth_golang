package main

import "fmt"

func main() {
	//a := 5
	//b := 8.1
	//fmt.Println(a + int(b))

	a := [5]int{1, 2, 3, 4, 5}
	t := a[2:4:5]
	fmt.Println(len(t), cap(t))
}
