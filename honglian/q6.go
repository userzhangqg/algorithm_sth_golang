package main

import "fmt"

/*
已作答：error,2,1,0,John,

正确答案：error,3,3,3,John,

知识点：defer参数快照
*/

func main() {
	Defer("John")
}

func Defer(name string) {
	defer func(param string) {
		fmt.Printf("%s,", param)
	}(name)
	for j := 0; j < 3; j++ {
		defer func() {
			fmt.Printf("%d,", j)
		}()
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%s,", err)
		}
	}()
	name = "Lee"
	panic("error")
	defer func() {
		fmt.Printf("end,")
	}()
}
